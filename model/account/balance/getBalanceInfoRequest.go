package balance

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
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
	return fmt.Sprintf("%sBalanceService/getBalanceInfo", model.BASE_URL_SMS)
}
