package InputMode

const (


	/**
 * 手输
 */
	HAND = "01"
/**
 * 刷卡，磁条卡
 */
SWIPE = "02"
/**
 * 扫码
 */
QRCODE = "03"
/**
 * Qpboc或简易流程
 */
QPBOC_OR_SIMPLE = "07"
/**
 * 标准PBOC借/贷记IC卡读入（非接触式）
 */
STANDARD_RF = "98"
/**
 * 芯片卡,插卡
 */
STANDARD_IC = "05"
/**
 * 集成电路卡，卡信息不可靠
 */
CHIP = "95"
/**
 * 采用非接触方式读取CUPMobile移动支付中的集成在手机中的芯片卡
 */
MOBILE_CHIP = "96"
	/**
	 * 非接触式磁条读入（MSD）
	 */
MSD = "91"
)