package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// userId - 账户ID
// balance - 账户余额
// pcBalance - 框架-基准资金包余额
// budget - 账户预算
// budgetType - 预算类型
// budgetOfflineTime - 到达预算下线时段
// cost - 账户累积消费
// excludeIp - IP排除列表
// openDomains - 账户开放域名列表
// payment - 账户投资总额
// regDomain - 账户注册域名
// regionTarget - 推广地域列表
// userStat - 用户状态
// userLevel - 客户权益-账户等级
// regionPriceFactor - 分地域出价系数
type GetAccountInfoRequest struct {
	AccountFields []string `json:"accountFields,omitempty"` // 指定需要返回的属性;
}

func (r GetAccountInfoRequest) Url() string {
	return fmt.Sprintf("%sAccountService/getAccountInfo", model.BASE_URL_SMS)
}
