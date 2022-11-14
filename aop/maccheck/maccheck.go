package maccheck

import (
	"NPP/byteutil"
	"NPP/easy8583"
	"NPP/logUtils"
	"NPP/lowcode/iso8583/model/terminal"
	"encoding/json"
	"strings"
)

func CheckMac(data8583 *easy8583.Easy8583, beforeHeadSize int) bool {

	// 1.根据终端号获取密钥   up.Field_R[40].Data
	terminalNo := string(data8583.Field_R[40].Data) //byteutil.BytesToHexString(up.Field_R[40].Data)
	logUtils.GlogInfo("终端号：" + terminalNo)
	isExit, err, terminals := terminal.QueryTerminalBySn(terminalNo);
	if isExit {
		jsonBytes, _ := json.Marshal(terminals[0])
		jsonStr := string(jsonBytes)
		logUtils.GlogInfo("terminalData：" + jsonStr)
	} else {
		logUtils.GlogInfo("QueryTerminalBySn err：" + err.Error())
		logUtils.GlogInfo("mac check err :there is not terminal record for this terminalId ")
		return false
	}

	mackey := byteutil.HexStringToBytes(terminals[0].MacKey)

	macData, err := easy8583.CalcMac16Byte(data8583.Txbuf[beforeHeadSize:], len(data8583.Txbuf)-beforeHeadSize-8, mackey)
	if err != nil {
		logUtils.GlogInfo("calc mac err：" + err.Error())
	}
	logUtils.GlogInfo("calc mac：" + byteutil.BytesToHexString(macData))
	logUtils.GlogInfo("data mac：" + byteutil.BytesToHexString(data8583.Field_R[63].Data))

	if strings.EqualFold(byteutil.BytesToHexString(macData), byteutil.BytesToHexString(data8583.Field_R[63].Data)) {
		return true
	} else {
		return false
	}

}