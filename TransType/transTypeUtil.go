package TransType

import (
	"NPP/TransType/InputMode"
	"NPP/TransType/MsgType"
	"NPP/TransType/ProcessType"
	"strings"
)

// 定义枚举类型 长度类型定义
const (
	/**
	 * 余额查询
	 */
	TRANS_BALANCE = 0
	/**
	 * 消费
	 */
	TRANS_SALE = 1
	/**
	 * 消费撤销
	 */
	TRANS_VOID_SALE = 2
	/**
	 * 退货
	 */
	TRANS_REFUND = 3
	/**
	 * 预授权
	 */
	TRANS_PREAUTH = 4
	/**
	 * 授权完成请求
	 */
	TRANS_AUTHSALE = 5

	/**
	 * 预授权撤销
	 */
	TRANS_VOID_PREAUTH = 6
	/**
	 * 预授权完成撤销
	 */
	TRANS_VOID_AUTHSALE = 7
	/**
	 * 签到
	 */
	TRANS_LOGIN = 8

	/**
	 * 签退
	 */
	TRANS_LOGOUT = 9

	/**
	 * 冲正
	 */
	TRANS_REVERSAL = 10

	/**
	 * 批结算
	 */
	TRANS_SETTLE = 11

	/**
	 * 批上送
	 */
	TRANS_BATCHUP = 12

	/**
	 * 状态上送
	 */
	TRANS_STATUS_SEND = 13
	/**
	 * 参数传递
	 */
	TRANS_PARAM_TRANSFER = 14

	/**
	 * AID状态上送
	 */
	TRANS_AID_STATUS = 15

	/**
	 * AID下载
	 */
	TRANS_LOAD_AID = 16

	/**
	 * AID end
	 */
	TRANS_LOAD_AID_END = 17

	/**
	 * CAPK状态上送
	 */
	TRANS_CAPK_STATUS = 18
	/**
	 * CAPK下载
	 */
	TRANS_LOAD_CAPK = 19

	/**
	 * CAPK end
	 */
	TRANS_LOAD_CAPK_END = 20

	/**
	 * 银联扫码
	 */
	TRANS_UNION_SCAN_PAY = 21
	/**
	 * 银联撤销
	 */
	TRANS_UNION_SCAN_VOID = 22
	/**
	 * 银联退货
	 */
	TRANS_UNION_SCAN_REFUND = 23
)

func AnTransType(msgtype string, processCodef03 string, inputModef22 string,
	serviceCodef25 string, field60_1 string, netManCodef60_2 string) int {
	var transType int = -1
	switch msgtype {
	case MsgType.PRE_AUTH:
		if ProcessType.PRE_AUTH == processCodef03 && "06" == serviceCodef25 {
			//预授权
			return TRANS_PREAUTH
		}
		if ProcessType.RETURNS == processCodef03 && "06" == serviceCodef25 {
			//预授权撤销
			return TRANS_VOID_PREAUTH
		}
		break
	case MsgType.FINANCE:
		if ProcessType.BALANCE == processCodef03 && "00" == serviceCodef25 {
			//余额查询
			return TRANS_BALANCE
		}
		if ProcessType.PAY == processCodef03 && "00" == serviceCodef25 {
			if len(inputModef22) > 0 && strings.HasPrefix(inputModef22, InputMode.QRCODE) {
				//银联扫码
				return TRANS_UNION_SCAN_PAY
			}
			//消费
			return TRANS_SALE
		}
		if ProcessType.RETURNS == processCodef03 && "00" == serviceCodef25 {
			if len(inputModef22) > 0 && strings.HasPrefix(inputModef22, InputMode.QRCODE) {
				//银联扫码撤销
				return TRANS_UNION_SCAN_VOID
			}
			//消费撤销
			return TRANS_VOID_SALE
		}
		if ProcessType.PAY == processCodef03 && "06" == serviceCodef25 {
			//授权完成
			return TRANS_AUTHSALE
		}
		if ProcessType.RETURNS == processCodef03 && "06" == serviceCodef25 {
			//授权完成撤销
			return TRANS_VOID_AUTHSALE
		}
		break
	case MsgType.REFUND:
		if ProcessType.RETURNS == processCodef03 && "00" == serviceCodef25 {
			if len(inputModef22) > 0 && strings.HasPrefix(inputModef22, InputMode.QRCODE) {
				//银联扫码退货
				return TRANS_UNION_SCAN_REFUND
			}
			//退货
			return TRANS_REFUND
		}
		break
	case MsgType.BATCH_UP:
		return TRANS_BATCHUP
	case MsgType.REVERSAL:
		return TRANS_REVERSAL
	case MsgType.SETTLE:
		return TRANS_SETTLE
	case MsgType.SCRIPT:
	case MsgType.LOGIN_OR_PARAMS:
		if len(netManCodef60_2) > 0 {
			if "001" == netManCodef60_2 || "003" == netManCodef60_2 || "004" == netManCodef60_2 || "005" == netManCodef60_2 || "006" == netManCodef60_2 {
				//签到
				return TRANS_LOGIN
			} else if "360" == netManCodef60_2 {
				//参数传递
				return TRANS_PARAM_TRANSFER
			} else if "380" == netManCodef60_2 {
				//下载AID
				return TRANS_LOAD_AID
			} else if "381" == netManCodef60_2 {
				//AID下载结束
				return TRANS_LOAD_AID_END
			} else if "370" == netManCodef60_2 {
				//下载CAPK
				return TRANS_LOAD_CAPK
			} else if "371" == netManCodef60_2 {
				//CAPK下载结束
				return TRANS_LOAD_CAPK_END
			}
		}
		break
	case MsgType.LOGOUT_OR_STATUS:
		if len(netManCodef60_2) > 0 {
			if "002" == netManCodef60_2 {
				//签退
				return TRANS_LOGOUT
			} else if "362" == netManCodef60_2 {
				//状态上送
				return TRANS_STATUS_SEND
			} else if "382" == netManCodef60_2 {
				//AID状态上送
				return TRANS_AID_STATUS
			} else if "372" == netManCodef60_2 || "373" == netManCodef60_2 {
				//CAPK状态上送
				return TRANS_CAPK_STATUS
			}
		}
		break
	default:
	}
	return transType
}
