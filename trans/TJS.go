package trans

import (
	"NPP/byteutil"
	"NPP/easy8583"
	"time"
)

/*
消费mms
*/
func ZJS8583(down *easy8583.Easy8583, up *easy8583.Easy8583) []byte {

	s := down
	field := down.Field_S

	s.Init8583Fields(field)

	// 0500/0510
	s.Msgtype[0] = 0x05
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

	//14域 卡有效期 N4
	field[13].Ihave = true
	field[13].Len = 2
	field[13].Data = []byte(d13)

	//15域 清算日期 N4
	field[14].Ihave = true
	field[14].Len = 2
	field[14].Data = byteutil.HexStringToBytes(time.Now().Format("0102"))

	//32域  标识码 N..11(LLVAR-BCD)，2个字节的长度值+最大11个字节的受理方标识码
	field[31].Ihave = true
	sqm := "01031000"
	field[31].Len = len(sqm)
	field[31].Data = byteutil.HexStringToBytes(sqm)

	//37域 "303436383739303837353634" 参考号 AN12，12个字节的定长字符域
	field[36].Ihave = true
	field[36].Data = byteutil.HexStringToBytes("303436383739303837353634")

	////39域 应答码 AN2，2个字节的定长字符域
	//field[38].Ihave = true
	//field[38].Data= byteutil.HexStringToBytes("3030")

	//41域，终端标识码 ANS8，8个字节的定长的字母、数字和特殊字符
	field[40].Ihave = true
	field[40].Len = up.Field_R[40].Len
	field[40].Data = up.Field_R[40].Data

	//42域，商户号  受卡方标识码 ANS15，15个字节30位的定长的字母、数字和特殊字符
	field[41].Ihave = true
	field[41].Len = 15
	field[41].Data = byteutil.HexStringToBytes("333038333530313933393930303830")

	//48域 (LLLVAR-BCD) 附加数据
	field[47].Ihave = true
	field[47].Len = up.Field_R[47].Len
	field[47].Data = up.Field_R[47].Data

	//49域，交易货币代码
	field[48].Ihave = true
	field[48].Data = byteutil.HexStringToBytes("313536")

	//60域 (LLLVAR-BCD) 8421 消息类型码/批次号/网络管理码
	field[59].Ihave = true
	field[59].Len = up.Field_R[59].Len //001100   1248
	field[59].Data = up.Field_R[59].Data

	//63域 LLLVAR  自定义域 0003435550
	field[62].Ihave = true
	field[62].Len = up.Field_R[62].Len
	field[62].Data = up.Field_R[62].Data

	//MAC，64域
	field[63].Ihave = true
	field[63].Len = 0x08
	field[63].Data = make([]byte, 8)

	/*报文组帧，自动组织这些域到Pack的TxBuffer中*/
	s.Pack8583Fields()
	return s.Txbuf
}
