package balance

// GetPaymentHistoryResponse 查询待加款信息 API Response
type GetPaymentHistoryResponse struct {
	Data []PaymentHistory `json:"data,omitempty"`
}

// PaymentHistory
type PaymentHistory struct {
	// Id 流水号
	Id int64 `json:"id,omitempty"`
	// UserId 用户id
	UserId int64 `json:"userId,omitempty"`
	// UserName 用户名
	UserName string `json:"userName,omitempty"`
	// MainAccountId 主账户id
	MainAccountId int64 `json:"mainAccountId,omitempty"`
	// MainAccountName 主账户名
	MainAccountName string `json:"mainAccountName,omitempty"`
	// Fund 金额，单位：元，正数表示加款，负数表示退款
	Fund float64 `json:"fund,omitempty"`
	// AccountName 账户类型
	AccountName string `json:"accountName,omitempty"`
	// FundPurpose 资金用途
	FundPurpose string `json:"fundPurpose,omitempty"`
	// PayMethodName 支付方式
	PayMethodName string `json:"payMethodName,omitempty"`
	// PayTime 加款时间，格式：yyyy-MM-dd HH:mm:ss
	PayTime string `json:"payTime,omitempty"`
	// OrderId 订单号
	OrderId int64 `json:"orderId,omitempty"`
}
