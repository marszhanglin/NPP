package sdtcp

import (
	"NPP/byteutil"
	"NPP/routers"
	"encoding/hex"
	"log"
	"net"
	"time"
)

func Run() {
	// lsof -i tcp:port    lsof -i tcp:12308
	//addr := "localhost:12308" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	addr := ":30000" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
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
	downmsg := routers.Rt(recvstr)
	log.Printf("send=====================================================\n")
	log.Printf("send len:[%d] data:[%s]\n", len(downmsg), byteutil.BytesToHexStr(downmsg, len(downmsg)))
	log.Printf("=========================================================\n")
	// 3.发送报文并关闭连接
	writeConn(conn, downmsg)
	// 4.
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
	log.Println("start to read from conn")
	n, err := conn.Read(lengthdata)
	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(lengthdata))
	}
	var receiveLength = ((int(lengthdata[0]) << 10) | int(lengthdata[1]))
	//receiveLength=11
	var buf = make([]byte, receiveLength)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read error:", err.Error())
			return ""
		} else {
			log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(buf)) //bytes-->16进制

		}
		break
	}

	recvstr := hex.EncodeToString(lengthdata) + hex.EncodeToString(buf)
	return recvstr
}
