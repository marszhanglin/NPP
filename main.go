package main

import (
	"NPP/cfg"
	"NPP/dao"
	"NPP/logUtils"
	"NPP/routershttp"
	"NPP/routersocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e *echo.Echo
)

func main() {

	//ipek, _ := hex.DecodeString("A1A3B4E3FD9B9CB480D432BB1B1A6FAF")
	//ksn, _ := hex.DecodeString("FF01A987654321000005")
	//// 明
	////plainText := []byte("0200303805802001060000000000000000050000001605350510140071000000353431333531323030363233333438383944323430353230313030303030383435303030004430303130313032323035323337373937303032303031313030363030313030313030303131303131303031310040303031323834304330303030303030303031373530303137383430433030303030303030303030330194df790656312e332e33df781430303030303030304e43393730303030313134359f0607a00000000310109f3901079f4005ff80f0a0019f02060000000005009f03060000000000009f26085d0a03685f7dd29b4f07a0000000031010820220009f36026c349f2701809f3303e0f8c89f1a0208409f350122950500000000005f2a0208409a032210149b0200009f21030534589c01009f3704c26743f55f2d047a68656e5f3401008407a00000000310109f1007060c0a03a000009f6e0420700000")
	//// 密文
	////cypherText, _ := hex.DecodeString("0AE3A50A54950F39A0069E550229B240AD89C9F9D65242773D25DBB13B45841860AB2369439D0F98A9A95E145C3C2A45385B6013F8F41BAFF0A59868761BEBA864AB585C6B6B212FC9B4815D7F2DD340025BB5786E32F365977C31764D4E765BA295AC6F3E20F3BC3A7107E14F1926442434349B042AAD70696510ED327C975BC3A40036C3DB8E8C29BF1EFF46DAF02CF358DB0E86ACAB9E72A300C6F9E7965BEB9E88CAFDEE90B56F1E5695172F13FACD6820DF904394A170AE4CA5D62C83159FAD1C0685DA6BE1352794898DE3A035DF2706AAEDD8650BE53B30CA0F467B7878E79D4F32B2FDA97A6EE5C9DB54956A6F9F0BFD064A70E14330E0A59D1597478BA6A9C1C08CAF79014F89E6414B4C8E8210792F8A192D98C9AA3C26E97E93A7602F17AD837E262DEF8D24B1D46A12597D0A089BDEFD2C19D0E9BFC4F3750058B85651A87C513ECCFEF44074937BB5D8B73689D3DF4FCF9262C30C6276B79EAB")
	//ksn, _ = hex.DecodeString("ff01a98765432100024d")
	//cypherText, _ := hex.DecodeString("7dfa462fb937652a66b245c7d87db8b372e88cce2d37f5a6e2ac34d4834879a48649bbc06092d5db6415c0470a22efa5425a6ba3a25361240dfd15b2d70e2165a9601d422d75278132a184b40383a0c22cf9441fa16e71655990c7931fc2d8da5c488fce455e2bfaa1de5d2a26fcac19ee6cdb2bcab6d5cd0b6559576a8c02d063dd86f783ab3179ad1a4a59440e7363e0675dad8752b9ca8e31e2e4ae10b9a6107bc543ea8542f708af0dd1a0723e13ecee71959cba5832d6eee62ce4f9dee8ed82e6e617beedfa2e81e480b4dda38707b874c323337d633e668b6a42352652ac9b3fd29b8302759d9b525d1990fde839c1b0ddef4c4dac6322ae682f0b3657390f34227d7daec01b2fd41c933410bc33de42769db52d875b705f7125209449cfb8e22e20c2f248f8d1961b6680fca67dc260057669b9f502c371a09c46716caedf044a249d044e00b4727354de55c2fa4d96d30292677eb4c24babd42fee2b2bcac6a20ef201d2bee5cb6d01071bb9aa669d9344ae0c42c27a62f975adf248")
	//go dukpttool.CBCDecrypterByIpek(ipek, ksn, cypherText)
	//
	//time.Sleep(5 * time.Second)

	logUtils.InitLog()

	//go dataans.Rt("01be600000000030303030303030304e373733303135343030323201807dfa462fb937652a66b245c7d87db8b372e88cce2d37f5a6e2ac34d4834879a48649bbc06092d5db6415c0470a22efa5425a6ba3a25361240dfd15b2d70e2165a9601d422d75278132a184b40383a0c22cf9441fa16e71655990c7931fc2d8da5c488fce455e2bfaa1de5d2a26fcac19ee6cdb2bcab6d5cd0b6559576a8c02d063dd86f783ab3179ad1a4a59440e7363e0675dad8752b9ca8e31e2e4ae10b9a6107bc543ea8542f708af0dd1a0723e13ecee71959cba5832d6eee62ce4f9dee8ed82e6e617beedfa2e81e480b4dda38707b874c323337d633e668b6a42352652ac9b3fd29b8302759d9b525d1990fde839c1b0ddef4c4dac6322ae682f0b3657390f34227d7daec01b2fd41c933410bc33de42769db52d875b705f7125209449cfb8e22e20c2f248f8d1961b6680fca67dc260057669b9f502c371a09c46716caedf044a249d044e00b4727354de55c2fa4d96d30292677eb4c24babd42fee2b2bcac6a20ef201d2bee5cb6d01071bb9aa669d9344ae0c42c27a62f975adf2480100020220202064322e31326666303161393837363534333231303030323464773030")

	//初始化数据库
	dao.DEBUG = true
	//dao.Db = dao.InitDB()

	//"./sdtcp"
	//go sdtcp.Run()

	cfg.InitViper()

	go routersocket.Run()

	// http
	initEcho()
	routershttp.EchoParamsRoute(e)
	go e.Logger.Fatal(e.Start(":4430"))

	//go
	//recvstr:="0060600601000061320032020505000020000000c18012000086303931303135313533303833353031393339393030383000620000000000150050000000000140040000000000000000000000000002002031353600110000000120100003303120"
	//"./routers"
	//log.Print(hex.EncodeToString(routers.Rt(recvstr)))

}

func initEcho() {
	//e := initEcho()
	e = echo.New()
	// Middleware
	e.Use(middleware.Logger()) //打印到控制台
	//打印到文件
	e.Use(middleware.Recover())
}