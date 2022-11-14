package routersocket

import (
	"NPP/logUtils"
	"NPP/routersocket/dataans"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Run() {
	// lsof -i tcp:port    lsof -i tcp:12308l
	//addr := "localhost:12308" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	addr := ":7890" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		go handle_conn(conn) //开启多个协程。
	}
}

func handle_conn(conn net.Conn) { //这个是在处理客户端会阻塞的代码。

	// 1.读取报文
	recvstr := readConn(conn)
	log.Printf("receive=====================================================\n")
	log.Printf("receive len:[%d] data:[%s]\n", len(recvstr), recvstr)
	log.Printf("============================================================\n")
	// 2.处理报文
	posreq := dataans.Pos2Cloud(recvstr)
	log.Printf("send=====================================================\n")
	//log.Printf("send len:[%d] data:[%s]\n", len(downmsg), byteutil.BytesToHexStr(downmsg, len(downmsg)))
	posreq.PrintFields()
	log.Printf("=========================================================\n")
	// 3.打包后台并向后台发送数据
	send2Clouddatas := make([][]byte, 3)    //length header(without sn) cleardata
	send2Clouddatas[1] = posreq.Header[0:5] // tpdu 去除sn
	send2Clouddatas[2] = posreq.ClearData
	//send2Clouddatas[2], _ = hex.DecodeString(Fill(DecHex(int64(len(send2Clouddatas[3]))), "0", 4, true))  // data不要长度
	send2Clouddatas[0], _ = hex.DecodeString(Fill(DecHex(int64(len(send2Clouddatas[1])+len(send2Clouddatas[2]))), "0", 4, true))
	send2cd := bytes.Join(send2Clouddatas, []byte(""))
	logUtils.Println("send2cd:" + hex.EncodeToString(send2cd))
	receive4cd := SendSocket("114.116.242.254:10002", send2cd)
	logUtils.Println("receive4cd:" + hex.EncodeToString(receive4cd))
	//4.打包并向终端返回
	cr := dataans.Cloud2Pos(posreq, receive4cd)
	recevice4Cloud := make([][]byte, 5) //length  header dataLength data footer
	recevice4Cloud[1] = cr.Header
	recevice4Cloud[2], _ = hex.DecodeString(Fill(DecHex(int64(len(cr.ClearData))), "0", 4, true))
	recevice4Cloud[3] = cr.ClearData
	recevice4Cloud[4], _ = hex.DecodeString("00" + cr.Footer[2:len(cr.Footer)]) // 返回的数据都是要不加密
	recevice4Cloud[0], _ = hex.DecodeString(Fill(DecHex(int64(len(recevice4Cloud[1])+len(recevice4Cloud[2])+len(recevice4Cloud[3])+len(recevice4Cloud[4]))), "0", 4, true))
	back2posdata := bytes.Join(recevice4Cloud, []byte(""))
	logUtils.Println("back2posdata:" + hex.EncodeToString(back2posdata))
	//5.发送报文并关闭连接
	writeConn(conn, back2posdata)
	//6.
	closeConn(conn)

}

// 发送数据
func writeConn(conn net.Conn, downmsg []byte) {
	//conn.Write(byteutil.HexStringToBytes(writeStr)) //通过conn的wirte方法将这些数据返回给客户端。
	conn.Write(downmsg) //通过conn的wirte方法将这些数据返回给客户端。

}

// 关闭连接
func closeConn(conn net.Conn) {
	time.Sleep(time.Minute)
	conn.Close() //与客户端断开连接。
}

// 接收数据
func readConn(conn net.Conn) string {
	// read from the connection
	var lengthdata = make([]byte, 2)
	logUtils.Println("start to read from conn")
	n, err := conn.Read(lengthdata)
	if err != nil {
		logUtils.Println("conn read error:" + err.Error())
	} else {
		log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(lengthdata))
	}
	//var receiveLength = BytesToInt(lengthdata)
	var receiveLength = ((int(lengthdata[0]) << 8) | int(lengthdata[1]))
	//receiveLength=11
	var buf = make([]byte, receiveLength)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			logUtils.Println("conn read error:" + err.Error())
			return ""
		} else {
			log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(buf)) //bytes-->16进制

		}
		break
	}

	recvstr := hex.EncodeToString(lengthdata) + hex.EncodeToString(buf)
	return recvstr
}

func SendSocket(address string, data []byte) []byte {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	_, err = conn.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	redata := readCloudConn(conn, false)
	//fmt.Fprintf(os.Stderr, "Fatal finish: %s", redata)
	conn.Close()
	return redata
}

// 接收数据
func readCloudConn(conn net.Conn, needLength bool) []byte {
	// read from the connection
	var lengthdata = make([]byte, 2)
	logUtils.Println("start to read from conn")
	n, err := conn.Read(lengthdata)
	if err != nil {
		logUtils.Println("conn read error:" + err.Error())
	} else {
		log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(lengthdata))
	}
	var receiveLength = ((int(lengthdata[0]) << 8) | int(lengthdata[1]))
	//receiveLength=11
	var buf = make([]byte, receiveLength)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			logUtils.Println("conn read error:" + err.Error())
			return []byte("")
		} else {
			log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(buf)) //bytes-->16进制

		}
		break
	}

	if needLength {
		lenData := make([][]byte, 2)
		lenData[0] = lengthdata
		lenData[1] = buf
		returnData := bytes.Join(lenData, []byte(""))
		return returnData
	} else {
		return buf
	}
}

func DecHex(n int64) string {
	if n < 0 {
		logUtils.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

func Fill(sour string, fillStr string, size int, isLeft bool) string {
	if len(sour) == 0 {
		sour = ""
	}

	fillLen := size - len(sour)
	fill := ""

	for i := 0; i < fillLen; i = i + 1 {
		fill = fill + fillStr
	}
	if isLeft {
		return fill + sour
	} else {
		return sour + fill
	}
}
