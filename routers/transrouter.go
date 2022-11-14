package routers

import (
	"NPP/TransType"
	"NPP/byteutil"
	"NPP/easy8583"
	"NPP/futils"
	"NPP/logUtils"
	"NPP/trans"
	"encoding/hex"
)

var (
	ManNum  string = "000000000000000"
	PosNum  string = "00000000"
	MainKey string = "00000000000000000000000000000000"
	TPDU    string = "6002800000"

	CommSn     int    = 1
	RecSn      int    = 1
	PiciNum    []byte = make([]byte, 3)
	LicenceNum        = []byte{0x33, 0x30, 0x36, 0x30}

	MacKey string = "0000000000000000" //加密校验码
)

func Rt(recvstr string) []byte {
	up8583 := easy8583.New8583() // 创建8583对象
	down8583 := easy8583.New8583()
	// 1.解析报文
	recv := byteutil.HexStringToBytes(recvstr)   //16进制字符串转bytes
	ret := up8583.Ans8583Fields(recv, len(recv)) // 解析8583报文
	if ret == 0 {
		logUtils.GlogInfo("解析成功")
		up8583.PrintFields(up8583.Field_R)
	} else {
		logUtils.GlogInfo("解析失败")
	}

	// 2.判断请求类型   根据：交易要素表
	// 消息类型
	msgtype := futils.BytesToHexString(up8583.Msgtype)
	// 处理码
	processCodef03 := "000000"
	inputModef22 := ""
	serviceCodef25 := "00"
	field60_1 := "00"
	netManCodef60_2 := ""
	if up8583.Field_R[3-1].Ihave {
		processCodef03 = hex.EncodeToString(up8583.Field_R[3-1].Data)
	}
	if up8583.Field_R[22-1].Ihave {
		inputModef22 = hex.EncodeToString(up8583.Field_R[22-1].Data)
	}
	if up8583.Field_R[25-1].Ihave {
		serviceCodef25 = hex.EncodeToString(up8583.Field_R[25-1].Data)
	}
	if up8583.Field_R[60-1].Ihave {
		field60_1 = hex.EncodeToString(up8583.Field_R[60-1].Data)
		field60_1 = string([]rune(field60_1)[:2])
		if (up8583.Field_R[60-1].Len >= 11) {
			netManCodef60_2 = string([]rune(hex.EncodeToString(up8583.Field_R[60-1].Data))[8:11])
		}
	}

	var downmsg []byte
	transType := TransType.AnTransType(msgtype, processCodef03, inputModef22, serviceCodef25, field60_1, netManCodef60_2)

	logUtils.GlogInfo("AnTransType:" + string(transType))

	switch transType {
	case TransType.TRANS_BALANCE:
	case TransType.TRANS_SALE:
		downmsg = trans.ZXF8583(down8583, up8583)
	case TransType.TRANS_VOID_SALE:
		downmsg = trans.ZXFCX8583(down8583, up8583)
	case TransType.TRANS_REFUND:
		downmsg = trans.ZTH8583(down8583, up8583)
	case TransType.TRANS_PREAUTH:
		downmsg = trans.ZYSQ8583(down8583, up8583)
	case TransType.TRANS_AUTHSALE:
		downmsg = trans.ZYSQWC8583(down8583, up8583)
	case TransType.TRANS_VOID_PREAUTH:
		downmsg = trans.ZYSQCX8583(down8583, up8583)
	case TransType.TRANS_VOID_AUTHSALE:
		downmsg = trans.ZYSQWCCX8583(down8583, up8583)
	case TransType.TRANS_LOGIN:
		downmsg = trans.ZQD8583(down8583, up8583)
	case TransType.TRANS_LOGOUT:
		downmsg = trans.ZQT8583(down8583, up8583)
	case TransType.TRANS_REVERSAL:
		downmsg = trans.ZReversal8583(down8583, up8583)
	case TransType.TRANS_SETTLE:
		downmsg = trans.ZJS8583(down8583, up8583)
	case TransType.TRANS_BATCHUP:
	case TransType.TRANS_STATUS_SEND:
	case TransType.TRANS_PARAM_TRANSFER:
	case TransType.TRANS_AID_STATUS:
		downmsg = trans.ZAIDStatus8583(down8583, up8583)
	case TransType.TRANS_LOAD_AID:
		downmsg = trans.ZLoadAID8583(down8583, up8583)
	case TransType.TRANS_LOAD_AID_END:
		downmsg = trans.ZLoadAIDEnd8583(down8583, up8583)
	case TransType.TRANS_CAPK_STATUS:
		downmsg = trans.ZCAPKStatus8583(down8583, up8583)
	case TransType.TRANS_LOAD_CAPK:
		downmsg = trans.ZLoadCAPK88583(down8583, up8583)
	case TransType.TRANS_LOAD_CAPK_END:
		downmsg = trans.ZLoadCAPKEnd88583(down8583, up8583)
	case TransType.TRANS_UNION_SCAN_PAY:
	case TransType.TRANS_UNION_SCAN_VOID:
	case TransType.TRANS_UNION_SCAN_REFUND:
	default:
		downmsg = []byte("未知消息类型")
	}

	return []byte(downmsg)

}

