package posreq

import (
	"NPP/byteutil"
	"NPP/logUtils"
	"encoding/hex"
	"fmt"
	"strings"
)

type PosReq struct {
	Len       []byte
	Header    []byte
	Data      []byte
	ClearData []byte
	DataLen   []byte
	Footer    string
	Ksn       []byte
}

type CloudResp struct {
	Header    []byte
	Data      []byte
	ClearData []byte
	Footer    string
}

/*
构造函数，初始化
len + header + data（Encrypt out） +  Footer
*/
func NewPosReq() *PosReq {
	var pr = new(PosReq)

	pr.Len = []byte{0x00, 0x00}

	pr.Header = make([]byte, 25)
	pr.DataLen = []byte{0x00, 0x00}
	pr.Ksn = make([]byte, 20)

	return pr
}

func NewCloudResp() *CloudResp {
	var pr = new(CloudResp)

	return pr
}

/*
解包终端请求报文
*/
func (pr *PosReq) UnpackPosReq(rxbuf string) int {
	pr.Len, _ = hex.DecodeString(rxbuf[0 : 2*2]) //报文长度 从0字节开始拷贝2字节
	pr.Header, _ = hex.DecodeString(rxbuf[2*2 : (25+2)*2])
	pr.DataLen, _ = hex.DecodeString(rxbuf[((25 + 2) * 2):((25 + 2 + 2) * 2)])
	var dataLen = len(rxbuf)/2 - (2 + 2 + 25 + 35) // len2 + header25 + data（Encrypt out） +  Footer35
	pr.Data = make([]byte, dataLen)
	pr.Data, _ = hex.DecodeString(rxbuf[((25 + 2 + 2) * 2):((25 + 2 + 2 + dataLen) * 2)])
	pr.Footer = rxbuf[((25 + 2 + 2 + dataLen) * 2):((25 + 2 + 2 + dataLen + 35) * 2)]
	pr.Ksn, _ = hex.DecodeString(pr.Footer[12*2 : (12+20)*2])
	pr.Ksn, _ = hex.DecodeString(string(pr.Ksn))
	pr.PrintFields()
	return 0
}

/*
打包平台返回的报文
*/
func (pr *PosReq) PackCloudResp(data4Cloud []byte, rxbuf string) int {
	pr.Len, _ = hex.DecodeString(rxbuf[0 : 2*2]) //报文长度 从0字节开始拷贝2字节
	pr.Header, _ = hex.DecodeString(rxbuf[2*2 : (25+2)*2])
	pr.DataLen, _ = hex.DecodeString(rxbuf[((25 + 2) * 2):((25 + 2 + 2) * 2)])
	var dataLen = len(rxbuf)/2 - (2 + 2 + 25 + 35) // len2 + header25 + data（Encrypt out） +  Footer35
	pr.Data = make([]byte, dataLen)
	pr.Data, _ = hex.DecodeString(rxbuf[((25 + 2 + 2) * 2):((25 + 2 + 2 + dataLen) * 2)])
	pr.Footer = rxbuf[((25 + 2 + 2 + dataLen) * 2):((25 + 2 + 2 + dataLen + 35) * 2)]
	pr.Ksn, _ = hex.DecodeString(pr.Footer[12*2 : (12+20)*2])
	pr.Ksn, _ = hex.DecodeString(string(pr.Ksn))
	pr.PrintFields()
	return 0
}

func Bcd2Number(bcd []byte) string {
	var number string
	for _, i := range bcd {
		number += fmt.Sprintf("%02X", i)
	}
	pos := strings.LastIndex(number, "F")
	if pos == 8 {
		return "0"
	}
	return number[pos+1:]
}

//func (pr *PosReq) UnpackPosReq(rxbuf []byte, rxlen int) int {
//	memcpy(pr.Len, rxbuf[0:], 2)     //报文长度 从0字节开始拷贝2字节
//	memcpy(pr.Header, rxbuf[2:], 25) //报文长度 从0字节开始拷贝2字节
//	memcpy(pr.DataLen, rxbuf[2+25:], 2)
//	var dataLen = rxlen - (2 + 2 + 25 + 35) // len2 + header25 + data（Encrypt out） +  Footer35
//	pr.Data = make([]byte, dataLen)
//	memcpy(pr.Data, rxbuf[(2+2+25):], dataLen)
//	memcpy(pr.Footer, rxbuf[(2+2+25+dataLen):], 35)
//	memcpy(pr.Ksn, pr.Footer[12:], 20)
//	pr.Txbuf = rxbuf
//	pr.PrintFields()
//	return 0
//}

func memcpy(dst, src []byte, size int) {
	for i := 0; i < size; i++ {
		dst[i] = src[i]
	}
	return
}

/*
打印信息，调试用
*/
func (pr *PosReq) PrintFields() {
	logUtils.Println("Print fields...")
	logUtils.Println("\n==========================================\n")
	logUtils.Println("Len:" + byteutil.BytesToHexString(pr.Len))
	logUtils.Println("Header:" + byteutil.BytesToHexString(pr.Header))
	logUtils.Println("DataLen:" + byteutil.BytesToHexString(pr.DataLen))
	logUtils.Println("Data:" + byteutil.BytesToHexString(pr.Data))
	logUtils.Println("ClearData:" + byteutil.BytesToHexString(pr.ClearData))
	logUtils.Println("Footer:" + pr.Footer)
	logUtils.Println("Ksn:" + byteutil.BytesToHexString(pr.Ksn))
	logUtils.Println("\n==========================================\n")
}
