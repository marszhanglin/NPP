package dukpttool

import (
	"NPP/byteutil"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	dukpt2 "github.com/wagner-aos/cryptokit/soft/dukpt"
	"log"
)

func CBCDecrypterByIpek(ipek []byte, ksn []byte, cypherText []byte) ([]byte, error) {

	log.Println("ipek: %x", byteutil.BytesToHexStr(ipek, len(ipek)))
	log.Println("ksn: %x", byteutil.BytesToHexStr(ksn, len(ksn)))
	log.Println("cypherText: %x", byteutil.BytesToHexStr(cypherText, len(cypherText)))

	error, dataKey := Ipek2DataKey(ipek, ksn)

	log.Println("dataKey: %x", byteutil.BytesToHexStr(dataKey, len(dataKey)))
	plainText := make([]byte, len(cypherText))
	plainText = tDesDecrypt(cypherText, dataKey)
	log.Println("PlainText: %x", byteutil.BytesToHexStr(plainText, len(plainText)))

	return plainText, error
}

func Ipek2DataKey(ipek []byte, ksn []byte) (error, []byte) {
	pek, error := dukpt2.DerivePekFromIpek(ipek, ksn)
	pekXor, _ := hex.DecodeString("00000000000000FF00000000000000FF")
	safeXORBytes(pek, pek, pekXor)
	log.Println("PEK: %x", byteutil.BytesToHexStr(pek, len(pek)))
	//// pek--mac密钥
	//macXor, _ := hex.DecodeString("000000000000FF00000000000000FF00")
	//// pek--pin密钥
	//pinXor, _ := hex.DecodeString("00000000000000FF00000000000000FF")
	// pek--data密钥
	dataXor, _ := hex.DecodeString("0000000000FF00000000000000FF0000")
	dataPek := make([]byte, len(dataXor))
	safeXORBytes(dataPek, pek, dataXor)
	// data--> dataDes
	dataDesML := make([]byte, 8)
	dataDesMR := make([]byte, 8)
	byteutil.Memcpy(dataDesML, dataPek, 8)
	byteutil.Memcpy(dataDesMR, dataPek[8:], 8)
	log.Println("dataPek: %x", byteutil.BytesToHexStr(dataPek, len(dataPek)))
	log.Println("dataDesML: %x", byteutil.BytesToHexStr(dataDesML, len(dataDesML)))
	log.Println("dataDesMR: %x", byteutil.BytesToHexStr(dataDesMR, len(dataDesMR)))
	dataDesML = tDesEncry(dataDesML, dataPek)
	dataDesMR = tDesEncry(dataDesMR, dataPek)
	dataKey := make([]byte, len(dataXor))
	byteutil.Memcpyse(dataKey, 0, dataDesML, 8)
	byteutil.Memcpyse(dataKey, 8, dataDesMR, 8)
	return error, dataKey
}

func safeXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func tDesEncry(data, key []byte) []byte {
	log.Println("TriDes---------------------------------s")
	tdes, _ := des.NewTripleDESCipher(BuildTdesKey(key))
	// 3des encrypt
	cbcEncrypt := cipher.NewCBCEncrypter(tdes, make([]byte, 8))
	resultEncrypt := make([]byte, len(data))
	cbcEncrypt.CryptBlocks(resultEncrypt, data)
	log.Println("resultEncrypt:", byteutil.BytesToHexStr(resultEncrypt, len(resultEncrypt)))
	log.Println("TriDes---------------------------------e")
	return resultEncrypt
}

func tDesDecrypt(data, key []byte) []byte {
	log.Println("TriDes---------------------------------s")
	tdes, _ := des.NewTripleDESCipher(BuildTdesKey(key))
	// 3des encrypt
	cbcEncrypt := cipher.NewCBCDecrypter(tdes, make([]byte, 8))
	resultEncrypt := make([]byte, len(data))
	cbcEncrypt.CryptBlocks(resultEncrypt, data)
	log.Println("resultEncrypt:", byteutil.BytesToHexStr(resultEncrypt, len(resultEncrypt)))
	log.Println("TriDes---------------------------------e")
	return resultEncrypt
}

func BuildTdesKey(key []byte) []byte {
	var finalKey []byte

	if len(key) == 24 {
		finalKey = key
	} else if len(key) == 16 {
		finalKey = make([]byte, 24)
		copy(finalKey, key)
		copy(finalKey[16:], key[:8])
	}

	return finalKey
}
