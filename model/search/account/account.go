package account

import "github.com/bububa/baidu-marketing/model"

// Account 账户信息
type Account struct {
	// UserId 账户ID
	UserId int64 `json:"userId,omitempty"`
	// Balance 账户余额
	Balance   float64 `json:"balance,omitempty"`
	PcBalance float64 `json:"pcBalance,omitempty"`
	// PcBalance 框架-基准资金包余额
	// Cost 账户累积消费，历史字段不建议参考，以数据报告为准
	Cost float64 `json:"cost,omitempty"`
	// Payment 账户投资总额，历史字段不建议参考，以数据报告为准
	Payment float64 `json:"payment,omitempty"`
	// Budget 账户预算
	Budget float64 `json:"budget,omitempty"`
	// RegionTarget 推广地域列表
	RegionTarget []int `json:"regionTarget,omitempty"`
	// ExcludeIps IP排除列表
	ExcludeIps []string `json:"excludeIps,omitempty"`
	// OpenDomains 账户开放域名列表(开放域名是在已有注册域名的基础上，再添加用户可以推广的域名，系统默认用户只能推广注册域名，物料URL需要和注册域名及开放域名的主域一致即可)
	OpenDomains []string `json:"openDomains,omitempty"`
	// BudgetType 预算类型; 0 - 不设置预算;1 - 日预算
	BudgetType *int `json:"budgetType,omitempty"`
	// RegDomain 账户注册域名
	RegDomain string `json:"regDomain,omitempty"`
	// UserStat 用户状态;1 - 开户金未到;2 - 正常生效;3 - 余额为零;4 - 未通过审核; 6 - 审核中; 7 - 被禁用; 11 - 预算不足
	UserStat int `json:"userStat,omitempty"`
	// BudgetOfflineTime 到达预算下线时段; 数组元素个数限制：最近有过下线时段的7个自然日的下线和上线时段（这7个自然日中若某日期距当前已超过30天，则不返回）；null：无到达预算下线时段；注：时间为date类型，格式示例”Jul 10, 2015 11:00:00 AM”
	BudgetOfflineTime []model.OfflineTime `json:"budgetOfflineTime,omitempty"`
	// UserLevel 客户权益-账户等级;1 - 三星客户;2 - 二星客户;3 - 一星客户;4 - 未生效客户
	UserLevel int `json:"userLevel,omitempty"`
	// RegionPriceFactor 分地域出价系数
	RegionPriceFactor []model.RegionPriceFactor `json:"regionPriceFactor,omitempty"`
}
