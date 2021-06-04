package balance

type GetPaymentHistoryResponse struct {
	Data []PaymentHistory `json:"data,omitempty"`
}

type PaymentHistory struct {
	Id              int64   `json:"id,omitempty"`              // 流水号
	UserId          int64   `json:"userId,omitempty"`          // 用户id
	UserName        string  `json:"userName,omitempty"`        // 用户名
	MainAccountId   int64   `json:"mainAccountId,omitempty"`   // 主账户id
	MainAccountName string  `json:"mainAccountName,omitempty"` // 主账户名
	Fund            float64 `json:"fund,omitempty"`            // 金额，单位：元，正数表示加款，负数表示退款
	AccountName     string  `json:"accountName,omitempty"`     // 账户类型
	FundPurpose     string  `json:"fundPurpose,omitempty"`     // 资金用途
	PayMethodName   string  `json:"payMethodName,omitempty"`   // 支付方式
	PayTime         string  `json:"payTime,omitempty"`         // 加款时间，格式：yyyy-MM-dd HH:mm:ss
	OrderId         int64   `json:"orderId,omitempty"`         // 订单号
}
