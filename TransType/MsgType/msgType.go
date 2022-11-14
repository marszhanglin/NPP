package MsgType

const (
	/**预授权类*/
	PRE_AUTH = "0100"
	/**金融类*/
	FINANCE = "0200"
	/**金融通知类--退货*/
	REFUND = "0220"
	/**批上送*/
	BATCH_UP = "0320"
	/**冲正*/
	REVERSAL = "0400"
	/**结算*/
	SETTLE = "0500"
	/**脚本通知*/
	SCRIPT = "0620"
	/**签到或参数下载类*/
	LOGIN_OR_PARAMS = "0800"
	/**签退或状态上送*/
	LOGOUT_OR_STATUS = "0820"
)
