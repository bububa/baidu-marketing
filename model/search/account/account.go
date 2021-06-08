package account

import "github.com/bububa/baidu-marketing/model"

type Account struct {
	UserId            int64                     `json:"userId,omitempty"`            // 账户ID
	Balance           float64                   `json:"balance,omitempty"`           // 账户余额
	PcBalance         float64                   `json:"pcBalance,omitempty"`         // 框架-基准资金包余额
	Cost              float64                   `json:"cost,omitempty"`              // 账户累积消费，历史字段不建议参考，以数据报告为准
	Payment           float64                   `json:"payment,omitempty"`           // 账户投资总额，历史字段不建议参考，以数据报告为准
	Budget            float64                   `json:"budget,omitempty"`            // 账户预算
	RegionTarget      []int                     `json:"regionTarget,omitempty"`      // 推广地域列表
	ExcludeIps        []string                  `json:"excludeIps,omitempty"`        // IP排除列表
	OpenDomains       []string                  `json:"openDomains,omitempty"`       // 账户开放域名列表(开放域名是在已有注册域名的基础上，再添加用户可以推广的域名，系统默认用户只能推广注册域名，物料URL需要和注册域名及开放域名的主域一致即可)
	BudgetType        *int                      `json:"budgetType,omitempty"`        // 预算类型; 0 - 不设置预算;1 - 日预算
	RegDomain         string                    `json:"regDomain,omitempty"`         // 账户注册域名
	UserStat          int                       `json:"userStat,omitempty"`          // 用户状态;1 - 开户金未到;2 - 正常生效;3 - 余额为零;4 - 未通过审核; 6 - 审核中; 7 - 被禁用; 11 - 预算不足
	BudgetOfflineTime []model.OfflineTime       `json:"budgetOfflineTime,omitempty"` // 到达预算下线时段; 数组元素个数限制：最近有过下线时段的7个自然日的下线和上线时段（这7个自然日中若某日期距当前已超过30天，则不返回）；null：无到达预算下线时段；注：时间为date类型，格式示例”Jul 10, 2015 11:00:00 AM”
	UserLevel         int                       `json:"userLevel,omitempty"`         // 客户权益-账户等级;1 - 三星客户;2 - 二星客户;3 - 一星客户;4 - 未生效客户
	RegionPriceFactor []model.RegionPriceFactor `json:"regionPriceFactor,omitempty"` // 分地域出价系数
}
