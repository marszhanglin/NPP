// Dbtool.go
package logUtils

import (
	"flag"
	"github.com/marszhanglin/logz"
	"log"
)

func Println(str string) {
	logz := logz.GetIns()
	logz.SugarInfo(str)
	log.Println(str)
	//logz:=	&logz{}
	//mLogZ := logz.InitCore(time.Now().Format("2006_0102_1504_05"))
	//mLogZ.SugarInfo(str)
}

// glog
func InitLog() {
	//  直接初始化，主要使服务器启动后自己直接加载，并不用命令行执行对应的参数
	flag.Set("alsologtostderr", "false") // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", "./logs")        // 日志文件保存目录
	flag.Set("v", "1")                   // 配置V输出的等级
	flag.Parse()                         // 1  解析命令行参数  go build ./hello_3_log.go   hello_3_log.exe -log_dir="./"
	//glog.Flush()                         // 4

}

//	func GlogInfo(str string) {
//		go fmt.Println(str)
//		go glog.Info(str)
//	}
//func GlogFlush() {
//	glog.Flush()
//}
