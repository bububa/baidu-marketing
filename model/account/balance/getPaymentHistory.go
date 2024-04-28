package balance

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetPaymentHistoryRequest 查询待加款信息
type GetPaymentHistoryRequest struct {
	// FundType 资金类型，必填，21（ka 现金），22（ka 优惠）
	FundType int `json:"fundType,omitempty"`
	// Id 记录id，第一页数据不填，第二页及往后的数据必须填，升序填当前页最大id值，降序填当前页最小id值，当前只支持降序
	Id uint64 `json:"id,omitempty"`
	// PageSize 每页数量大小，必填，每页数量必须大于0，小于等于1000
	PageSize int64 `json:"pageSize,omitempty"`
	// StartTime 支付时间-开始时间，选填，格式：yyyy-MM-dd HH:mm:ss
	StartTime string `json:"startTime,omitempty"`
	// EndTime 支付时间-结束时间，选填，格式：yyyy-MM-dd HH:mm:ss
	EndTime string `json:"endTime,omitempty"`
}

func (r GetPaymentHistoryRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "PaymentFeedService/getPaymentHistory")
}

// GetPaymentHistoryResponse 查询待加款信息 API Response
type GetPaymentHistoryResponse struct {
	Data []PaymentHistory `json:"data,omitempty"`
}

// PaymentHistory
type PaymentHistory struct {
	// Id 流水号
	Id uint64 `json:"id,omitempty"`
	// UserId 用户id
	UserId uint64 `json:"userId,omitempty"`
	// UserName 用户名
	UserName string `json:"userName,omitempty"`
	// MainAccountId 主账户id
	MainAccountId uint64 `json:"mainAccountId,omitempty"`
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
	OrderId uint64 `json:"orderId,omitempty"`
}
