package trans

//import "encoding/hex"

import (
	"NPP/byteutil"
	"NPP/desutil"
	"NPP/easy8583"
	"NPP/logUtils"
	"NPP/lowcode/iso8583/model/terminal"
	"encoding/json"
	"strings"
	"time"
)

/*
银联8583 交易组包
*/
func ZQD8583(down *easy8583.Easy8583, up *easy8583.Easy8583) []byte {

	// 1.根据终端号获取密钥   up.Field_R[40].Data
	terminalNo := string(up.Field_R[40].Data) //byteutil.BytesToHexString(up.Field_R[40].Data)
	logUtils.GlogInfo("终端号：" + terminalNo)
	isExit, err, terminals := terminal.QueryTerminalBySn(terminalNo)
	if isExit {
		jsonBytes, _ := json.Marshal(terminals[0])
		jsonStr := string(jsonBytes)
		logUtils.GlogInfo("terminalData：" + jsonStr)
	} else {
		logUtils.GlogInfo("QueryTerminalBySn err：" + err.Error())
		return nil
	}
	//

	s := down
	field := down.Field_S

	s.Init8583Fields(field)

	//签到类型 0810
	s.Msgtype[0] = 0x08
	s.Msgtype[1] = 0x10

	//11域 系统跟踪号（流水号）   N6
	field[10].Ihave = true
	field[10].Len = 3
	field[10].Data = up.Field_R[10].Data

	//"092309" mmddHH
	//12域 时间 N6
	field[11].Ihave = true
	field[11].Len = 3
	time12 := time.Now().Format("150405")
	d12 := byteutil.HexStringToBytes(time12)
	field[11].Data = []byte(d12)

	//13域 日期 N4
	field[12].Ihave = true
	field[12].Len = 2
	date13 := time.Now().Format("0102")
	d13 := byteutil.HexStringToBytes(date13)
	field[12].Data = []byte(d13)

	//32域  标识码 N..11(LLVAR)，2个字节的长度值+最大11个字节的受理方标识码
	field[31].Ihave = true
	sqm := "01031000"
	field[31].Len = len(sqm)
	field[31].Data = byteutil.HexStringToBytes(sqm)

	//37域 "303436383739303837353634" 参考号 AN12，12个字节的定长字符域
	field[36].Ihave = true
	field[36].Data = byteutil.HexStringToBytes("303436383739303837353634")

	//39域 应答码 AN2，2个字节的定长字符域
	field[38].Ihave = true
	field[38].Data = byteutil.HexStringToBytes("3030")

	//41域，终端标识码 ANS8，8个字节的定长的字母、数字和特殊字符
	field[40].Ihave = true
	field[40].Len = up.Field_R[40].Len
	field[40].Data = up.Field_R[40].Data

	//42域，商户号  受卡方标识码 ANS15，15个字节30位的定长的字母、数字和特殊字符
	field[41].Ihave = true
	field[41].Len = 15
	field[41].Data = byteutil.HexStringToBytes("333038333530313933393930303830")

	//60域 (LLLVAR-BCD 消息类型码/批次号/网络管理码
	field[59].Ihave = true
	field[59].Len = 0x12                                       //001100   1248
	field[59].Data = byteutil.HexStringToBytes("000000010040") //[]byte("0011000000010040")

	//62  LLLVAR 自定义域 3组工作密钥（ping密钥-密码密钥，mac密钥-报文密钥，磁道加密密钥） (32+8)*3=120
	pingMiKey, _ := desutil.Des3Encrypt(byteutil.HexStringToBytes(terminals[0].PinKey), byteutil.HexStringToBytes(terminals[0].MasterKey))
	pingMiKeyHex := byteutil.BytesToHexString(pingMiKey)
	logUtils.GlogInfo("pin密钥密文:" + pingMiKeyHex)
	pingMiKeyCKV, _ := desutil.Des3Encrypt(byteutil.X00, byteutil.HexStringToBytes(terminals[0].PinKey))
	pingMiKeyCKVHex := byteutil.BytesToHexString(pingMiKeyCKV)
	logUtils.GlogInfo("pin密钥CK:" + pingMiKeyCKVHex)

	macMiKey, _ := desutil.Des3Encrypt(byteutil.HexStringToBytes(terminals[0].MacKey), byteutil.HexStringToBytes(terminals[0].MasterKey))
	macMiKeyHex := byteutil.BytesToHexString(macMiKey)
	logUtils.GlogInfo("mac密钥密文:" + macMiKeyHex)
	macMiKeyCKV, _ := desutil.Des3Encrypt(byteutil.X00, byteutil.HexStringToBytes(terminals[0].MacKey))
	macMiKeyCKVHex := byteutil.BytesToHexString(macMiKeyCKV)
	logUtils.GlogInfo("mac密钥CK:" + macMiKeyCKVHex)

	trackMiKey, _ := desutil.Des3Encrypt(byteutil.HexStringToBytes(terminals[0].TrackKey), byteutil.HexStringToBytes(terminals[0].MasterKey))
	trackMiKeyHex := byteutil.BytesToHexString(trackMiKey)
	logUtils.GlogInfo("track密钥密文:" + trackMiKeyHex)
	trackMiKeyCKV, _ := desutil.Des3Encrypt(byteutil.X00, byteutil.HexStringToBytes(terminals[0].TrackKey))
	trackMiKeyCKVHex := byteutil.BytesToHexString(trackMiKeyCKV)
	logUtils.GlogInfo("track密钥CK:" + trackMiKeyCKVHex)

	f61 := strings.ToUpper(pingMiKeyHex + pingMiKeyCKVHex[0:8] + macMiKeyHex + macMiKeyCKVHex[0:8] + trackMiKeyHex + trackMiKeyCKVHex[0:8])
	logUtils.GlogInfo("62预报文:" + f61)

	field[61].Ihave = true
	field[61].Len = 0x60
	field[61].Data = byteutil.HexStringToBytes(f61)
	//field[61].Data = byteutil.HexStringToBytes("950973182317F80B950973182317F80B00962B60F679786E2411E3DE0000000000000000ADC67D84A0C45C59F1E549BBA0C45C59F1E549BBE2F24340")
	// 11111111111111111111111111111111(主密钥明文)加密22222222222222222222222222222222(工作密钥明文16字节32位)得95097318 2317F80B 95097318 2317F80B
	//22222222222222222222222222222222(工作密钥明文)加密0000000000000000(默认check值)得00962B60 AA556E65(checkValue,取前4字节8位)
	// 950973182317F80B950973182317F80B00962B60F679786E2411E3DE0000000000000000ADC67D84A0C45C59F1E549BBA0C45C59F1E549BBE2F24340
	// （16+4）*3=60
	//60域
	//field[59].Ihave = true
	//field[59].Len = 0x13
	//field[59].Data = make([]byte, 7)
	//field[59].Data[0] = 0x22
	//futils.Memcpy(field[59].Data[1:], PiciNum, 3)
	//field[59].Data[4] = 0x00
	//field[59].Data[5] = 0x06
	//field[59].Data[6] = 0x00

	/*报文组帧，自动组织这些域到Pack的TxBuffer中*/
	s.Pack8583Fields()
	down.Ans8583Fields(s.Txbuf, len(s.Txbuf))
	//log.Println("bbbbb"+hex.EncodeToString(s.Txbuf))
	return s.Txbuf
}

//60000006016132003202050810003800010AC00014000058092309021308010310003034363837393038373536343030303931303135313533303833353031393339393030383000110000000100400060950973182317F80B950973182317F80B00962B60F679786E2411E3DE0000000000000000ADC67D84A0C45C59F1E549BBA0C45C59F1E549BBE2F24340
//// 根据上送报文，获取签到报文
//func GetQD8583(up *easy8583.Easy8583) string {
//	qd := "008D"                                        //报文长度 	N2
//	qd = qd + "6000000601613200320205"                  //head
//	qd = qd + "0810"                                    //msgtype
//	qd = qd + "003800010AC00014"                        //bitmap   	N16  11,12,13,32,37,39,41,42,60,62
//	qd = qd + hex.EncodeToString(up.Field_R[11-1].Data) //11 系统跟踪号（流水号）   N6
//	qd = qd + "092309"                                  //12 时间 N6
//	qd = qd + "0213"                                    //13 日期 N4
//	qd = qd + "0801031000"                              //32 标识码 N..11(LLVAR)，2个字节的长度值+最大11个字节的受理方标识码
//	qd = qd + "303436383739303837353634"                //37 参考号 AN12，12个字节的定长字符域
//	qd = qd + "3030"                                    //39 应答码 AN2，2个字节的定长字符域
//	qd = qd + "3039313031353135"                        //41 终端标识码 ANS8，8个字节的定长的字母、数字和特殊字符
//	qd = qd + "333038333530313933393930303830"          //42 受卡方标识码 ANS15，15个字节的定长的字母、数字和特殊字符
//	qd = qd + "0011000000010040"                        //60 消息类型码/批次号/网络管理码
//	//62 自定义域 主密钥
//	qd = qd + "0060950973182317F80B950973182317F80B00962B60F679786E2411E3DE0000000000000000ADC67D84A0C45C59F1E549BBA0C45C59F1E549BBE2F24340"
//	return qd
//}
