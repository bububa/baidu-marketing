package account

import "github.com/xiaoshouchen/baidu-marketing/model"

// Account 账户信息
type Account struct {
	// UserId 账户ID
	UserId uint64 `json:"userId,omitempty"`
	// Cid 账户主体ID
	Cid uint64 `json:"cid,omitempty"`
	// LiceName 账户主体名称
	LiceName string `json:"liceName,omitempty"`
	// Balance 账户余额, 账户余额，对应财务平台"推广共享资金"字段
	Balance float64 `json:"balance,omitempty"`
	// PcBalance 框架-基准资金包余额, 仅部分直销框架客户需要关注（此部分客户的账户可用余额为pcBalance+balance）
	PcBalance float64 `json:"pcBalance,omitempty"`
	// Cost 账户累积消费，历史字段不建议参考，以数据报告为准
	Cost float64 `json:"cost,omitempty"`
	// Payment 账户投资总额，历史字段不建议参考，以数据报告为准
	Payment float64 `json:"payment,omitempty"`
	// Budget 账户预算, 当设置为日预算时，取值范围：[50, 10000000]； 当不设置预算时，输入任意值均默认为0
	Budget float64 `json:"budget,omitempty"`
	// RegionTarget 推广地域列表
	RegionTarget []int64 `json:"regionTarget,omitempty"`
	// ExcludeIps IP排除列表, 支持添加/删除ipv4和ipv6类型ip, 例如：2001:0DB8:0000:0023:0008:0800:200C:417A
	ExcludeIps []string `json:"excludeIps,omitempty"`
	// OpenDomains 账户开放域名列表(开放域名是在已有注册域名的基础上，再添加用户可以推广的域名，系统默认用户只能推广注册域名，物料URL需要和注册域名及开放域名的主域一致即可)
	OpenDomains []string `json:"openDomains,omitempty"`
	// BudgetType 预算类型; 0 - 不设置预算;1 - 日预算
	BudgetType int `json:"budgetType,omitempty"`
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
	// GeolocationStatus 推广地域地理位置选项 取值范围：枚举值，列表如下
	// 0 - 该地区内或搜索意图在该地区的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户、在搜索词中对该地区表现出明确兴趣的用户）
	// 1 - 该地区内的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户）
	// 原账户层级的“搜索意图地域词扩展”已升级为计划地域设置中的推广地域地理位置选项
	GeolocationStatus int `json:"geoLocationStatus,omitempty"`
	// ExcludeQueryRegionStatus 搜索意图地域词排除; 取值范围：枚举值，列表如下
	// true - 启用
	// false - 关闭
	ExcludeQueryRegionStatus bool `json:"excludeQueryRegionStatus,omitempty"`
	// LongMonitorSublink 创意组件-文字链监控代码; 可填充监控代码默认值："product=xinxizhenlie_longsublink"，如需自定义，请直接录入监控代码（字符限制255）
	LongMonitorSublink string `json:"longMonitorSublink,omitempty"`
	// AccountMonitorURL 通用点击监测地址; 当前通用点击检测地址长度限制1024
	AccountMonitorURL string `json:"accountMonitorURL,omitempty"`
}
