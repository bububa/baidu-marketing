package balance

// GetBalanceInfoResponse 查询账户余额成分 API Response
type GetBalanceInfoResponse struct {
	Data []BalanceInfo `json:"data,omitempty"`
}

// BalanceInfo 账户余额
type BalanceInfo struct {
	// Balance 余额
	Balance float64 `json:"balance,omitempty"`
	// Cash 现金余额
	Cash float64 `json:"cash,omitempty"`
	// BonusPure 优惠余额
	BonusPure float64 `json:"bonusPure,omitempty"`
	// Compen 补偿余额
	Compen float64 `json:"compen,omitempty"`
	// Bonus
	Bonus float64 `json:"bonus,omitempty"`
	// Invest
	Invest float64 `json:"invest,omitempty"`
	// Orderrow 订单行号，该字段仅KA客户返回
	Orderrow int64 `json:"orderrow,omitempty"`
}
