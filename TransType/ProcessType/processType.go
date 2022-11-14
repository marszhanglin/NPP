package ProcessType


const (
	/**
	* 扣款
	*/
	PAY = "000000"
	/**
	 * 预授权
	 */
	PRE_AUTH = "030000"
	/**
	 * 退款
	 */
	RETURNS = "200000"
	/**
	 * 查余
	 */
	BALANCE = "310000"
	/**
	 * 圈存
	 */
	LOAD = "600000"
	/**
	 * 非指定账户圈存
	 */
	NOT_APPOINTED_LOAD = "620000"
	/**
	 * 现金圈存
	 */
	CASH_SAVING = "630000"
	/**
	 * 现金撤销
	 */
	CASH_VOID = "170000"
)
