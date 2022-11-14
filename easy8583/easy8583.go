package easy8583

import (
	"NPP/desutil"
	"NPP/logUtils"
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Field struct {
	Ihave bool   //是否存在该域
	Ltype int    //长度类型 （NOVAR，LLVAR，LLLVAR）  如果是NOVAR定长，就需要配置Len
	Dtype int    //数据类型 （BCD,ASCII）
	Len   int    //域的数据内容的长度
	Data  []byte //域的有效数据
}

type Easy8583 struct {
	Len     []byte
	Tpdu    []byte
	Head    []byte
	Msgtype []byte
	Bitmap  []byte

	Txbuf []byte

	Field_S []Field //发送的域 定义 会去组发送的签名
	Field_R []Field //接收的域 定义 会去解接收的包名
}

// 定义枚举类型 长度类型定义
const (
	NOVAR  = iota //value = 0,定长,
	LLVAR         //value = 1，长度为1字节
	LLLVAR        //value = 2，长度为2字节

)

// 定义枚举类型 数据类型定义
const (
	UN  = iota //value = 0, 未定义，定长的域无需关注类型
	BIN        //value = 1，BIN
	BCD        //value = 2，BCD
)

var (
	PingKey = []byte{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22}
	MacKey  = []byte{0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33}
	CDKey   = []byte{0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44}
)

/*
设置工作秘钥,算MAC用
*/
func (ea *Easy8583) SetMacKey(strkey string) {

	MacKey = hexStringToBytes(strkey)

}

// 各个域的初始配置
func (ea *Easy8583) Init8583Fields(fds []Field) {

	for i := 0; i < 64; i++ {
		fds[i].Ihave = false
	}

	toZero(ea.Bitmap)

	fds[0].Ltype = 0

	fds[1].Ltype = LLVAR //LLVAR
	fds[1].Dtype = BCD

	fds[2].Ltype = NOVAR
	fds[2].Len = 3

	fds[3].Ltype = 0
	fds[3].Len = 6

	fds[4].Ltype = LLVAR
	fds[5].Ltype = LLVAR
	fds[6].Ltype = LLVAR
	fds[7].Ltype = LLVAR
	fds[8].Ltype = LLVAR
	fds[9].Ltype = LLVAR

	fds[10].Ltype = NOVAR
	fds[10].Len = 3

	fds[11].Ltype = NOVAR
	fds[11].Len = 3

	fds[12].Ltype = NOVAR
	fds[12].Len = 2

	fds[13].Ltype = NOVAR
	fds[13].Len = 2

	fds[14].Ltype = NOVAR
	fds[14].Len = 2

	fds[15].Ltype = LLVAR
	fds[16].Ltype = LLVAR
	fds[17].Ltype = LLVAR
	fds[18].Ltype = LLVAR
	fds[19].Ltype = LLVAR
	fds[20].Ltype = LLVAR

	fds[21].Ltype = NOVAR
	fds[21].Len = 2
	fds[22].Ltype = NOVAR
	fds[22].Len = 2
	fds[23].Ltype = NOVAR
	fds[23].Len = 2
	fds[24].Ltype = NOVAR
	fds[24].Len = 1
	fds[25].Ltype = NOVAR
	fds[25].Len = 1

	fds[26].Ltype = LLVAR
	fds[27].Ltype = LLVAR
	fds[28].Ltype = LLVAR
	fds[29].Ltype = LLVAR
	fds[30].Ltype = LLVAR

	fds[31].Ltype = LLVAR //LLVAR
	fds[31].Dtype = BCD

	fds[32].Ltype = LLVAR //LLVAR
	fds[33].Ltype = LLVAR //LLVAR
	fds[34].Ltype = LLVAR //LLVAR
	fds[34].Dtype = BCD
	fds[35].Ltype = LLVAR //LLVAR
	fds[35].Dtype = BCD

	fds[36].Ltype = NOVAR //AN12
	fds[36].Len = 6

	fds[37].Ltype = NOVAR
	fds[37].Len = 3
	fds[38].Ltype = NOVAR
	fds[38].Len = 1

	fds[39].Ltype = LLVAR

	fds[40].Ltype = NOVAR
	fds[40].Len = 8

	fds[41].Ltype = NOVAR
	fds[41].Len = 15

	fds[42].Ltype = NOVAR
	fds[42].Len = 20

	fds[43].Ltype = LLVAR

	fds[46].Ltype = LLLVAR

	fds[47].Ltype = LLLVAR
	fds[47].Dtype = BCD

	fds[48].Ltype = NOVAR
	fds[48].Len = 3
	fds[51].Ltype = NOVAR
	fds[51].Len = 8
	fds[52].Ltype = NOVAR
	fds[52].Len = 8

	fds[54].Ltype = LLLVAR //LLLVAR
	fds[54].Dtype = BIN

	fds[56].Ltype = LLLVAR //LLLVAR
	fds[56].Dtype = BIN

	fds[57].Ltype = LLLVAR
	fds[57].Dtype = BCD
	fds[58].Ltype = LLLVAR
	fds[58].Dtype = BCD
	fds[59].Ltype = LLLVAR
	fds[59].Dtype = BCD
	fds[60].Ltype = LLLVAR
	fds[60].Dtype = BCD

	fds[58].Ltype = LLLVAR

	fds[59].Ltype = LLLVAR
	fds[59].Dtype = BCD

	fds[60].Ltype = LLLVAR
	fds[60].Dtype = BCD

	fds[61].Ltype = LLLVAR
	fds[62].Ltype = LLLVAR

	fds[63].Ltype = NOVAR
	fds[63].Len = 8

}

/*
构造函数，初始化
*/

func New8583() *Easy8583 {

	var ea = new(Easy8583)
	ea.Txbuf = make([]byte, 0, 1024)
	ea.Txbuf = ea.Txbuf[0:17]

	ea.Len = []byte{0x00, 0x00}
	//ea.Tpdu = []byte{0x60, 0x05, 0x01, 0x00, 0x00}//[]byte{0x60, 0x05, 0x01, 0x00, 0x00} //
	//ea.Head = []byte{0x61, 0x31, 0x00, 0x31, 0x11, 0x08} //[]byte{0x61, 0x31, 0x00, 0x31, 0x11, 0x08}//6000000601613200320205
	ea.Tpdu, _ = hex.DecodeString("6006010000")   //[]byte{0x60, 0x05, 0x01, 0x00, 0x00} //
	ea.Head, _ = hex.DecodeString("613200320205") //[]byte{0x61, 0x31, 0x00, 0x31, 0x11, 0x08}//6000000601613200320205

	ea.Msgtype = []byte{0x08, 0x00}

	ea.Bitmap = make([]byte, 8)

	ea.Field_S = make([]Field, 64)
	ea.Field_R = make([]Field, 64)

	ea.Init8583Fields(ea.Field_S)
	ea.Init8583Fields(ea.Field_R)

	return ea
}

func memcpy(dst, src []byte, size int) {
	for i := 0; i < size; i++ {
		dst[i] = src[i]
	}
	return
}

func bytesToHexStr(data []byte, lenth int) string {
	buf := data[0:lenth]
	hexStr := fmt.Sprintf("%x", buf)
	//fmt.Println(hexStr)
	return hexStr

}

// bytes to hex string
func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// hex string to bytes
func hexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}

// 例：0x19 --> 19, 0x0119 -> 119
func bcdToInt(data []byte, lenth int) int {
	buf := data[0:lenth]
	hexStr := fmt.Sprintf("%x", buf)
	out, _ := strconv.ParseInt(hexStr, 10, 32)
	return int(out)

}
func toZero(p []byte) {
	for i := range p {
		p[i] = 0
	}
}

/*
计算银联8583通信MAC
*/
func dataXor1(src []byte, dest []byte, size int) {

	fmt.Println(hex.EncodeToString(src[0:8]), "---", hex.EncodeToString(dest))

	for i := 0; i < size; i++ {
		dest[i] ^= src[i]
	}

}

func dataXor(src []byte, dest []byte, size int, out []byte) {
	for i := 0; i < size; i++ {
		out[i] = dest[i] ^ src[i]
	}

}

/*
*

	计算8583报文mac，加密部分包含msge到63域，(不包含len tpdu head)
*/
func CalcMac16Byte(buf []byte, bufsize int, mackey []byte) ([]byte, error) {

	block := make([]byte, 1024)
	val := make([]byte, 8)
	memcpy(block, buf, bufsize)

	logUtils.GlogInfo(hex.EncodeToString(buf))

	x := bufsize / 8 //计算有多少个完整的块
	n := bufsize % 8

	if n != 0 {
		x += 1 //将补上的这一块加上去
	}
	j := 0
	for i := 0; i < x; i++ {
		dataXor1(block[j:], val, 8)
		fmt.Println("第", i, "组明文:", hex.EncodeToString(block[j:j+8]))
		fmt.Println("第", i, "组异或:", hex.EncodeToString(val))

		j += 8
	}

	Bbuf := fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X", val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7])
	//fmt.Printf("Bbuf:%s\n",Bbuf)
	Abuf := make([]byte, 8)
	mac, err := desutil.Des3Encrypt([]byte(Bbuf[0:8]), mackey)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("mac1:%x\n",bytesToHexString(mac))
	dataXor(mac, []byte(Bbuf[8:]), 8, Abuf)
	mac, err = desutil.Des3Encrypt(Abuf, mackey)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("mac2:%s\n",bytesToHexString(mac))
	outmac := fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5], mac[6], mac[7])
	//fmt.Printf("outmac:%s\n",outmac)
	return []byte(outmac[0:8]), nil

}

/*
*

	计算8583报文mac，加密部分包含msge到63域，(不包含len tpdu head)
*/
func CalcMac8Byte(buf []byte, bufsize int, mackey []byte) ([]byte, error) {

	block := make([]byte, 1024)
	val := make([]byte, 8)
	memcpy(block, buf, bufsize)

	x := bufsize / 8 //计算有多少个完整的块
	n := bufsize % 8

	if n != 0 {
		x += 1 //将补上的这一块加上去
	}
	j := 0
	for i := 0; i < x; i++ {
		dataXor1(block[j:], val, 8)
		j += 8
	}

	Bbuf := fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X", val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7])
	fmt.Printf("Bbuf:%s\n", Bbuf)
	Abuf := make([]byte, 8)
	fmt.Println(bytesToHexString([]byte(Bbuf[0:8])))
	mac, err := desutil.DesEncrypt([]byte(Bbuf[0:8]), mackey) //16字节的密钥，需要用3des计算，8字节的用单des就好
	if err != nil {
		return nil, err
	}
	//fmt.Printf("mac1:%x\n",mac)
	dataXor(mac, []byte(Bbuf[8:]), 8, Abuf)
	mac, err = desutil.DesEncrypt(Abuf, mackey)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("mac2:%x\n",mac)
	outmac := fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5], mac[6], mac[7])
	//fmt.Printf("outmac:%s\n",outmac)
	return []byte(outmac[0:8]), nil

}

/*
8583报文打包,args传入个工作秘钥
*/
func (ea *Easy8583) Pack8583Fields() int {
	fmt.Printf("pack 8583 fields\n")
	//ea.Txbuf[]
	ea.Txbuf = ea.Txbuf[0:17]
	toZero(ea.Txbuf)

	j := 0
	len := 17
	tmplen := 0
	seat := 0x80
	for i := 0; i < 64; i++ {
		//fmt.Println("F:",i)
		seat = (seat >> 1)
		if (i % 8) == 0 {
			j++
			seat = 0x80
		}
		if ea.Field_S[i].Ihave {
			ea.Bitmap[j-1] |= byte(seat)
			if ea.Field_S[i].Ltype == NOVAR {
				ea.Txbuf = ea.Txbuf[0 : len+ea.Field_S[i].Len]
				memcpy(ea.Txbuf[len:], ea.Field_S[i].Data, ea.Field_S[i].Len)
				len += ea.Field_S[i].Len

			} else if ea.Field_S[i].Ltype == LLVAR {
				ea.Txbuf = ea.Txbuf[0 : len+1]
				ea.Txbuf[len] = byte(ea.Field_S[i].Len)

				tmplen = bcdToInt(ea.Txbuf[len:], 1)
				if ea.Field_S[i].Dtype == BCD {
					tmplen = ((tmplen / 2) + (tmplen % 2))
				}
				len += 1
				ea.Txbuf = ea.Txbuf[0 : len+tmplen]
				memcpy(ea.Txbuf[len:], ea.Field_S[i].Data, tmplen)
				len += tmplen

			} else if ea.Field_S[i].Ltype == LLLVAR {
				ea.Txbuf = ea.Txbuf[0 : len+2]
				ea.Txbuf[len] = byte(ea.Field_S[i].Len >> 8)
				ea.Txbuf[len+1] = byte(ea.Field_S[i].Len)

				tmplen = bcdToInt(ea.Txbuf[len:], 2)
				if ea.Field_S[i].Dtype == BCD {
					tmplen = ((tmplen / 2) + (tmplen % 2))
				}
				len += 2
				ea.Txbuf = ea.Txbuf[0 : len+tmplen]
				memcpy(ea.Txbuf[len:], ea.Field_S[i].Data, tmplen)
				len += tmplen

			}

		}

	}

	//报文总长度
	ea.Txbuf[0] = byte((len - 2) >> 8)
	ea.Txbuf[1] = byte((len - 2))
	memcpy(ea.Len, ea.Txbuf, 2)
	memcpy(ea.Txbuf[2:], ea.Tpdu, 5)
	//memcpy(ea.Txbuf[7:], ea.Head, 6)  //  海外没有请求好
	memcpy(ea.Txbuf[7:], ea.Msgtype, 2)
	memcpy(ea.Txbuf[9:], ea.Bitmap, 8)
	//如果64域存在，自动计算MAC并填充
	if ea.Field_S[63].Ihave {
		//txbuf := []byte{0x00,0x69,0x60,0x01,0x38,0x00,0x00,0x61,0x31,0x00,0x31,0x11,0x08,0x02,0x00,0x30,0x20,0x04,0x80,0x00,0xc0,0x80,0x31,0x00,0x00,0x00,0x30,0x30,0x30,0x30,0x30,0x30,0x00,0x00,0x02,0x03,0x20,0x00,0x33,0x34,0x33,0x38,0x36,0x30,0x31,0x33,0x38,0x39,0x38,0x34,0x33,0x30,0x34,0x34,0x31,0x31,0x31,0x30,0x30,0x31,0x32,0x31,0x35,0x36,0x00,0x24,0x41,0x33,0x30,0x31,0x39,0x36,0x32,0x32,0x32,0x36,0x37,0x35,0x32,0x38,0x31,0x34,0x36,0x34,0x32,0x39,0x38,0x36,0x33,0x34,0x00,0x13,0x22,0x00,0x00,0x80,0x00,0x06,0x00}
		mac, err := CalcMac8Byte(ea.Txbuf[13:], len-13-8, MacKey)
		if err != nil {
			fmt.Println(err)
			panic("calc mac error!")
		}
		//fmt.Printf("mac:%x", mac)
		memcpy(ea.Field_S[63].Data, mac, 8)
		memcpy(ea.Txbuf[len-8:], mac, 8)
	}

	return 0
}

/*
8583报文解包
*/
func (ea *Easy8583) Ans8583Fields(rxbuf []byte, rxlen int) int {
	fmt.Printf("ans 8583 fields\n")
	ea.Init8583Fields(ea.Field_R)

	len := 0
	tmplen := 0
	bitMap := make([]byte, 8) //位图
	var seat, buf uint64 = 1, 0

	memcpy(ea.Len, rxbuf[0:], 2) //报文长度 从0字节开始拷贝2字节
	memcpy(ea.Tpdu, rxbuf[2:], 5)

	// 国内有消息头
	//memcpy(ea.Head, rxbuf[7:], 6)     //消息头
	//memcpy(ea.Msgtype, rxbuf[13:], 2) //消息类型
	//memcpy(ea.Bitmap, rxbuf[15:], 8)  //位图
	//memcpy(bitMap, rxbuf[15:], 8) //拷贝位图 从15字节开始拷贝8字节
	//len += 23

	// 海外没有消息头
	//memcpy(ea.Head, rxbuf[7:], 6)     //消息头
	memcpy(ea.Msgtype, rxbuf[7:], 2) //消息类型
	memcpy(ea.Bitmap, rxbuf[9:], 8)  //位图
	memcpy(bitMap, rxbuf[9:], 8)     //拷贝位图 消息类型后八个字节是位图
	len += 17

	for i := 0; i < 8; i++ { //算出位图
		buf = ((buf << 8) | uint64(bitMap[i]))
	}

	for i := 0; i < 64; i++ {
		if (buf & (seat << uint(63-i))) > 0 { //根据位图判断该域是否有值
			ea.Field_R[i].Ihave = true        //该域有值
			if ea.Field_R[i].Ltype == NOVAR { //判断该域的类型 ,为定长时
				ea.Field_R[i].Data = make([]byte, ea.Field_R[i].Len)
				memcpy(ea.Field_R[i].Data, rxbuf[len:], ea.Field_R[i].Len) //拷贝定长
				len += ea.Field_R[i].Len

			} else if ea.Field_R[i].Ltype == LLVAR { //判断该域的类型 ,为变长，长度位1字节

				ea.Field_R[i].Len = int(rxbuf[len]) //一个字节计算该域长度

				tmplen = bcdToInt(rxbuf[len:], 1)
				if ea.Field_R[i].Dtype == BCD {
					tmplen = ((tmplen / 2) + (tmplen % 2))
				}
				len += 1
				ea.Field_R[i].Data = make([]byte, tmplen)
				memcpy(ea.Field_R[i].Data, rxbuf[len:], tmplen)
				len += tmplen

			} else if ea.Field_R[i].Ltype == LLLVAR { //判断该域的类型 ,为变长，长度位2字节

				ea.Field_R[i].Len = ((int(rxbuf[len]) << 8) | int(rxbuf[len+1]))

				tmplen = bcdToInt(rxbuf[len:], 2)
				if ea.Field_R[i].Dtype == BCD {
					tmplen = ((tmplen / 2) + (tmplen % 2))
				}
				len += 2
				ea.Field_R[i].Data = make([]byte, tmplen)
				memcpy(ea.Field_R[i].Data, rxbuf[len:], tmplen)
				len += tmplen
				logUtils.GlogInfo("LLLVAR-i：" + string(i))

			}

		}

	}
	ea.Txbuf = rxbuf[0:len]
	logUtils.GlogInfo(hex.EncodeToString(ea.Txbuf))
	if len > rxlen {
		return 1
	}

	return 0
}

/*
打印信息，调试用
*/
func (ea *Easy8583) PrintFields(fds []Field) {
	fmt.Println("Print fields...")
	fmt.Printf("\n==========================================\n")
	fmt.Printf("Len:\t%s\n", bytesToHexString(ea.Len))
	fmt.Printf("Tpdu:\t%s\n", bytesToHexString(ea.Tpdu))
	fmt.Printf("Head:\t%s\n", bytesToHexString(ea.Head))
	fmt.Printf("Msge:\t%s\n", bytesToHexString(ea.Msgtype))
	fmt.Printf("Bitmap:\t%s\n", bytesToHexString(ea.Bitmap))
	fmt.Printf("\n==========================================\n")
	for i := 0; i < 64; i++ {
		if fds[i].Ihave {
			fmt.Printf("[field:%d] ", i+1)
			if fds[i].Ltype == LLVAR {
				fmt.Printf("[len:%02x] ", fds[i].Len)
			} else if fds[i].Ltype == LLLVAR {
				fmt.Printf("[len:%04x] ", fds[i].Len)
			}

			fmt.Printf("[%s]\n", bytesToHexString(fds[i].Data))
			fmt.Printf("\n------------------------------\n")

		}
	}
}

func main() {

	fmt.Println("test...")
}
