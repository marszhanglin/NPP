package trans

//import "encoding/hex"

import (
	"NPP/byteutil"
	"NPP/easy8583"
	"time"
)

/*
签退组包
*/
func ZQT8583(down *easy8583.Easy8583, up *easy8583.Easy8583) []byte {

	s := down
	field := down.Field_S

	s.Init8583Fields(field)

	//签到类型 0810
	s.Msgtype[0] = 0x08
	s.Msgtype[1] = 0x30

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

	//60域 (LLLVAR) 8421 消息类型码/批次号/网络管理码
	field[59].Ihave = true
	field[59].Len = 0x11 //001100   1248
	field[59].Data = byteutil.HexStringToBytes("000000010020")

	/*报文组帧，自动组织这些域到Pack的TxBuffer中*/
	s.Pack8583Fields()
	down.Ans8583Fields(s.Txbuf, len(s.Txbuf))
	return s.Txbuf
}
