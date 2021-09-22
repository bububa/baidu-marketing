package balance

type GetBalanceInfoResponse struct {
	Data []BalanceInfo `json:"data,omitempty"`
}

type BalanceInfo struct {
	Balance   float64 `json:"balance,omitempty"`   // 余额
	Cash      float64 `json:"cash,omitempty"`      // 现金余额
	BonusPure float64 `json:"bonusPure,omitempty"` // 优惠余额
	Compen    float64 `json:"compen,omitempty"`    // 补偿余额
	Bonus     float64 `json:"bonus,omitempty"`
	Invest    float64 `json:"invest,omitempty"`
	Orderrow  int64   `json:"orderrow,omitempty"` // 订单行号，该字段仅KA客户返回
}
