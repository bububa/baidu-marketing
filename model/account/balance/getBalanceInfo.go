package balance

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetBalanceInfoRequest 查询账户余额成分 API Request
// 1：推广共享资金（for KA客户&直销非框客户）
// 502：原生推广（for KA客户）
// 400：框架-基准资金包（for 直销框架客户）
// 430：框架-原生资金包（for直销框架客户）
type GetBalanceInfoRequest struct {
	ProductIds []int `json:"productIds,omitempty"` // 资金包ID
}

func (r GetBalanceInfoRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "BalanceService/getBalanceInfo")
}

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
	Orderrow uint64 `json:"orderrow,omitempty"`
}
