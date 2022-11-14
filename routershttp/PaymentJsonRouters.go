package routershttp

import (
	"NPP/cfg"
	"NPP/logUtils"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
)

var (
	projectName string
	validate    *validator.Validate
)

func EchoParamsRoute(e *echo.Echo) {
	validate = validator.New()
	e.POST("/1.0.0/param/", param)
	e.GET("/1.0.0/param/", param)
	e.GET("/flush", flushLog)
	//echoRouteSaveflow()
	//echoRouteUserManager()
}

func flushLog(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func param(c echo.Context) error {
	epi := c.QueryParam("epi")
	logUtils.Println("epi:" + epi)
	action := c.QueryParam("action")
	logUtils.Println("action:" + action)
	did := c.QueryParam("did")
	logUtils.Println("did:" + did)
	rspStr := ""
	switch action {
	case "REQUEST":
		//rspStr = "{\"app_download\":\"0\",\"param_download\":\"1\",\"app_package_id\":\"266\",\"current_app_version\":\"1.1.3\",\"last_param_success\":\"2022-09-08 08:45:20\",\"timezone\":\"EST\",\"time\":\"09-08-2022 04:49:35\",\"logo_checksum\":[\"6264bb2ecd614cb184f20c2f271565e8\",\"\"]}\n"
		rspStr = cfg.V.GetString("nnp.param.REQUEST")
		return c.String(http.StatusOK, rspStr)
		//return c.JSON(http.StatusOK, respone)
	case "DOWNLOAD":
		//rspStr = "{\"EMV_CONFIG\":{\"EMV_CONFIGURE\":[[\"E0B8C8\",\"E0F8C8\",\"F000F0A001\",\"0840\",\"0840\",\"USD\",\"22\",\"0\",\"2\",\"0840\",\"2\",\"000000\",\"2\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"0\",\"0\",\"9F3704\",\"9F02065F2A029A039C0195059F3704\",\"0\"]]},\"MENU\":{\"MENU_LIST\":[[\"1\",\"Payment\",\"0\",\"0\",\"\",\"0\"],[\"2\",\"Favorites\",\"0\",\"10\",\"\",\"0\"],[\"3\",\"Settlement\",\"0\",\"9\",\"1234\",\"0\"],[\"4\",\"Reprint\",\"0\",\"11\",\"1234\",\"0\"],[\"5\",\"Reports\",\"0\",\"12\",\"1234\",\"0\"],[\"6\",\"Utility\",\"0\",\"0\",\"1234\",\"0\"],[\"7\",\"HostUtility\",\"0\",\"0\",\"1234\",\"0\"],[\"8\",\"Setup\",\"0\",\"0\",\"\",\"0\"],[\"9\",\"RemoteDiagnostics\",\"0\",\"34\",\"1234\",\"0\"],[\"10\",\"Credit\",\"1\",\"0\",\"\",\"0\"],[\"11\",\"Debit\",\"1\",\"0\",\"\",\"0\"],[\"12\",\"EBTFood\",\"1\",\"0\",\"\",\"0\"],[\"13\",\"EBTCash\",\"1\",\"0\",\"\",\"0\"],[\"14\",\"Cash\",\"1\",\"0\",\"\",\"0\"],[\"15\",\"Sale\",\"10\",\"1\",\"\",\"1\"],[\"16\",\"Void\",\"10\",\"2\",\"1234\",\"1\"],[\"17\",\"Auth\",\"10\",\"3\",\"1234\",\"1\"],[\"18\",\"Ticket\",\"10\",\"4\",\"1234\",\"1\"],[\"19\",\"Refund\",\"10\",\"5\",\"1234\",\"1\"],[\"20\",\"Sale\",\"11\",\"1\",\"\",\"2\"],[\"21\",\"Refund\",\"11\",\"5\",\"1234\",\"2\"],[\"22\",\"Sale\",\"12\",\"1\",\"\",\"3\"],[\"23\",\"VoucherSale\",\"12\",\"6\",\"\",\"3\"],[\"24\",\"Refund\",\"12\",\"5\",\"1234\",\"3\"],[\"26\",\"FoodBalance\",\"12\",\"8\",\"\",\"3\"],[\"27\",\"Sale\",\"13\",\"1\",\"\",\"4\"],[\"28\",\"CashBalance\",\"13\",\"8\",\"\",\"4\"],[\"29\",\"Sale\",\"14\",\"1\",\"\",\"6\"],[\"30\",\"Refund\",\"14\",\"5\",\"\",\"6\"],[\"31\",\"SoftwareDownload\",\"6\",\"0\",\"\",\"0\"],[\"32\",\"CommConfig\",\"6\",\"0\",\"\",\"0\"],[\"33\",\"TMS\",\"31\",\"0\",\"\",\"0\"],[\"34\",\"Full\",\"33\",\"13\",\"\",\"0\"],[\"35\",\"Partial\",\"33\",\"14\",\"\",\"0\"],[\"36\",\"USB\",\"31\",\"15\",\"\",\"0\"],[\"37\",\"Media\",\"32\",\"16\",\"\",\"0\"],[\"38\",\"Config\",\"32\",\"17\",\"\",\"0\"],[\"39\",\"Void\",\"7\",\"18\",\"1234\",\"0\"],[\"40\",\"Pre-SaleTicket\",\"7\",\"19\",\"\",\"0\"],[\"41\",\"TipAdjust\",\"7\",\"20\",\"\",\"0\"],[\"42\",\"Ticket\",\"7\",\"21\",\"1234\",\"0\"],[\"45\",\"ClearReversal\",\"7\",\"24\",\"\",\"0\"],[\"46\",\"ResetTerminal\",\"7\",\"25\",\"\",\"0\"],[\"47\",\"CallMe\",\"7\",\"26\",\"\",\"0\"],[\"48\",\"ChangeDate&Time\",\"8\",\"27\",\"\",\"0\"],[\"50\",\"KeyboardBeep\",\"8\",\"29\",\"\",\"0\"],[\"51\",\"DisplayContrast\",\"8\",\"30\",\"\",\"0\"],[\"53\",\"ChangePassword\",\"8\",\"32\",\"\",\"0\"],[\"54\",\"TerminalReboot\",\"8\",\"33\",\"\",\"0\"],[\"55\",\"Withdrawal\",\"13\",\"44\",\"\",\"4\"],[\"56\",\"Debug\",\"8\",\"45\",\"\",\"0\"],[\"58\",\"RKI\",\"7\",\"46\",\"\",\"0\"],[\"59\",\"Gift\",\"1\",\"0\",\"\",\"0\"],[\"60\",\"Sale\",\"59\",\"1\",\"\",\"5\"],[\"61\",\"AddValue\",\"59\",\"51\",\"\",\"5\"],[\"62\",\"Balance\",\"59\",\"38\",\"\",\"5\"],[\"64\",\"Activate\",\"59\",\"49\",\"\",\"5\"],[\"64\",\"Deactivate\",\"59\",\"50\",\"1234\",\"5\"],[\"65\",\"Clerk\\/Server\",\"7\",\"28\",\"\",\"0\"],[\"67\",\"PRESALE\",\"10\",\"19\",\"\",\"1\"]]},\"GPRS\":{\"GPRS_TABLE\":[[\"30\",\"30\",\"30\",\"3\"]]},\"ETHERNET\":{\"ETH_TABLE\":[[\"30\",\"0\"]]},\"DIAL\":{\"DIAL_TABLE\":[[\"8666094332\",\"8666094332\",\"8666094332\",\"8666094332\",\"60\",\"0001\",\"0002\",\"0003\",\"0004\",\"00\",\"\",\"60\",\"60\",\"90\",\"\",\"\",\"\",\"\"]]},\"WIFI\":{\"WIFI_TABLE\":[[\"30\",\"0\"]]},\"TERM_CONFIG\":{\"TERM_CONFIGURE\":[[\"1931197138\",\"0\",\"EST\",\"1\",\"1\",\"1\",\"0\",\"30\",\"3\",\"1\",\"0\",\"0\",\"2\"]]},\"MER_INFO\":{\"MERCHANT_INFO\":[[\"1\",\"MerchantIndustry\"]]},\"TRAN_FLAG\":{\"TRANSACTION_FLAG\":[[\"7\",\"1\",\"1\",\"0\",\"1\",\"11000\",\"1\",\"1\",\"10000\",\"9999999\",\"1\",\"1\",\"1\",\"21\",\"24\",\"0\",\"1\",\"0\",\"000000003500\",\"0\",\"\",\"\",\"1\"]]},\"HEADER\":{\"HEADER_LIST\":[[\"NewlandMerchant\",\"234thstreet\",\"ASTORIANY11106\",\"222-222-2222\"]]},\"FOOTER\":{\"FOOTER_LIST\":[[\"ThankYou\"]]},\"DISCLAIMER\":{\"DISCLAIMER_LIST\":[[\"Cardholderacknowledges\",\"receiptofgoodsand\",\"obligationssetforth\",\"bythecardholder's\",\"agreement with issuer.\"]]},\"HW_CONFIG\":{\"HW_CONFIGURE\":[[\"80\",\"\",\"0\",\"80\",\"1\",\"0\",\"30\",\"1800\",\"300\",\"3600\"]]},\"SUPPORT\":{\"SUPPORT_NO\":[[\"\",\"\"]]},\"TIP\":{\"TIP_VALUES\":[[\"1\",\"015000\",\"020000\",\"022000\",\"025000\",\"00000\",\"0350000\"]]},\"PROMO_MSG\":{\"PROMOTIONAL_MESSAGE_CUSTOMER\":[[\"0\",\"12202019\",\"01012020\",\"\"],[\"0\",\"12202019\",\"01012020\",\"\"],[\"0\",\"12202019\",\"01012020\",\"\"]],\"PROMOTIONAL_MESSAGE_MERCHANT\":[[\"0\",\"12202019\",\"01012020\",\"\"]]},\"AUTO_TAX\":{\"AUTO_TAX_VALUES\":[[\"1\",\"State Tax\",\"00600\"]]},\"CUSTOM_FEE\":{\"CUSTOM_FEE_VALUES\":[[\"1\",\"2\",\"1\",\"1\",\"Non-Cash Charge\",\"999999\",\"00001\",\"00001\",\"1\"]]},\"LOGO\":{\"FLAG\":[[\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"]]},\"TERM_ADTL_INFO\":{\"TERM_ADTL_DATA\":[[\"CLERK\",\"000000001000\",\"0\",\"1\",\"1\",\"0\",\"0\",\"3\",\"1\",\"1\",\"1\",\"1\"]]},\"HOME_SCREEN_BIN_CONFIG_INFO\":{\"HOME_SCREEN_BIN_CONFIG_DATA\":[[\"0\",\"0\",\"2\",\"\",\"0\",\"1\",\"0\",\"0\",\"1\",\"0\",\"0\"]]},\"HOST_COMMS\":{\"TRAN\":[[\"txnuat.valorpaytech.com\",\"192.168.34.135\",\"192.168.34.135\",\"7890\"]],\"TMSAPP\":[[\"http://192.168.34.135:4430/1.0.0/app/?epi=\",\"000.000.000.000\",\"000.000.000.000\",\"0000\"]],\"TMSPAR\":[[\"http://192.168.34.135:4430/1.0.0/param/?epi=\",\"000.000.000.000\",\"000.000.000.000\",\"0000\"]],\"ERECPT\":[[\"http://192.168.34.135:4430/\",\"000.000.000.000\",\"000.000.000.000\",\"0000\"]],\"RD\":[[\"tms.isoaccess.com\",\"108.166.035.083\",\"108.166.035.083\",\"1883\"]],\"CONFIG\":[[\"3\",\"1\",\"30\",\"60\",\"90\"]],\"BO\":[[\"http://192.168.34.135:4430/heartbeat/HeartBeat.php\",\"000.000.000.000\",\"000.000.000.000\",\"0000\"]],\"RKI\":[[\"\",\"61.219.144.216\",\"61.219.144.216\",\"7020\"]],\"TMSVAS\":[[\"http://192.168.34.135:4430/1.0.0/param/tms/?epi=\",\"000.000.000.000\",\"000.000.000.000\",\"0000\"]]},\"TRAN_ADTL_INFO\":{\"TRAN_ADTL_VALUE\":[[\"1\",\"000000000100\",\"1\",\"1\",\"0\",\"0\",\"\",\"1\",\"0\",\"0\",\"\",\"\"]]}}\n"
		rspStr = cfg.V.GetString("nnp.param.DOWNLOAD")
		return c.String(http.StatusOK, rspStr)
		//return c.JSON(http.StatusOK, respone)
	case "DONE":
		//rspStr = "{\"error_no\":\"S00\",\"error_desc\":\"SUCCESS\"}"
		rspStr = cfg.V.GetString("nnp.param.DONE")
		return c.String(http.StatusOK, rspStr)
		//return c.JSON(http.StatusOK, respone)
	}

	return c.String(http.StatusOK, "unKnow action ")
	//return c.JSON(http.StatusOK, rsp)
}

func getFormParam(c echo.Context, key string) string {

	value := c.Request().PostForm[key]
	if value == nil {
		return ""
	} else {
		return value[0]
	}
}

func getHeaderParam(c echo.Context, key string) string {
	value := c.Request().Header.Get(key)
	return value
}

func getFormParamInt64(c echo.Context, key string) int64 {

	value := c.Request().PostForm[key]
	if value == nil {
		return 0
	} else {
		valueint64, _ := strconv.ParseInt(value[0], 10, 64)
		return valueint64
	}
}

func getResPonse(retCode string, retMsg string) *ResponseBody {
	response := &ResponseBody{}
	response.RetCode = retCode
	response.RetMsg = retMsg
	return response
}

type ResponseBody struct {
	RetCode string            `json:"retCode" xml:"retCode_"`
	Body    map[string]string `json:"body" xml:"body_"`
	RetMsg  string            `json:"retMsg" xml:"retMsg_"`
}

func getBaseResPonse(retCode string, retMsg string) *Response {
	response := &Response{}
	response.RetCode = retCode
	response.RetMsg = retMsg
	log.Printf(retCode + ":" + retMsg)
	return response
}

type ParamRspReq struct {
	RetCode string                 `json:"retCode" xml:"retCode_"`
	Body    map[string]interface{} `json:"body" xml:"body_"`
	RetMsg  string                 `json:"retMsg" xml:"retMsg_"`
}

type Response struct {
	RetCode string                 `json:"retCode" xml:"retCode_"`
	Body    map[string]interface{} `json:"body" xml:"body_"`
	RetMsg  string                 `json:"retMsg" xml:"retMsg_"`
}

//#string到int
//int,err:=strconv.Atoi(string)
//#string到int64
//int64, err := strconv.ParseInt(string, 10, 64)
//#int到string
//string:=strconv.Itoa(int)
//#int64到string
//string:=strconv.FormatInt(int64,10)
