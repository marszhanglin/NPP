package dataans

import (
	"NPP/byteutil"
	"NPP/dukpttool"
	"NPP/logUtils"
	"NPP/routersocket/dataans/posreq"
	"encoding/hex"
	"strings"
)

var (
	ipek []byte
)

// len + header + data（Encrypt out） +  Footer
func Pos2Cloud(recStr string) *posreq.PosReq {

	//recv, _ := hex.DecodeString(recStr) //16进制字符串转bytes

	posReq := posreq.NewPosReq()
	ret := posReq.UnpackPosReq(recStr)

	posReq.PrintFields()

	if strings.HasPrefix(posReq.Footer, "01") {
		logUtils.Println("需要解密")
		ipek, _ = hex.DecodeString("A1A3B4E3FD9B9CB480D432BB1B1A6FAF")
		clearData, _ := dukpttool.CBCDecrypterByIpek(ipek, posReq.Ksn, posReq.Data)
		posReq.ClearData = make([]byte, len(clearData))
		byteutil.Memcpyse(posReq.ClearData, 0, clearData, len(clearData))
		logUtils.Println("数据明文:" + byteutil.BytesToHexString(posReq.ClearData))
	} else {
		posReq.ClearData = posReq.Data
	}

	posReq.PrintFields()
	if ret == 0 {
		logUtils.Println("解析成功")
		//up8583.PrintFields(up8583.Field_R)
	} else {
		logUtils.Println("解析失败")
	}

	return posReq
}

func Cloud2Pos(p2cReq *posreq.PosReq, data []byte) *posreq.CloudResp {
	cr := posreq.NewCloudResp()
	cr.Header = p2cReq.Header        //  包含sn
	cr.ClearData = data[5:len(data)] //  不包含长度
	cr.Footer = p2cReq.Footer
	return cr
}
