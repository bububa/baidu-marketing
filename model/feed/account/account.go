package account

// Account 账户信息
type Account struct {
	// UserId 账户ID
	UserId uint64 `json:"userId,omitempty"`
	// Cid 账户主体ID
	// 客户不存在主体，则cid为0
	Cid uint64 `json:"cid,omitempty"`
	// Balance 账户余额。此字段仅用于查询，updateAccountFeed接口传此字段无效。
	Balance float64 `json:"balance,omitempty"`
	// Budget 账户预算。默认为0,表示不限预算。正常取值范围为[50-9999999.99]
	Budget float64 `json:"budget,omitempty"`
	// BalancePackage 资金包类型。此字段仅用于查询，updateAccountFeed接口传此字段无效。取值范围如下：0：原生资金包 1：凤巢资金包 2：代理商原生资金包
	BalancePackage int `json:"balancePackage,omitempty"`
	// UserStat 账户状态。此字段仅用于查询，updateAccountFeed接口传此字段无效。取值范围如下：1：开户金未到 2：生效 3：账户余额为0 4：被拒绝 6：审核中 7：被禁用 8：待激活 11：账户预算不足
	UserStat int `json:"userStat,omitempty"`
	// UaStat 是否开通feed产品线权限。此字段仅用于查询，updateAccountFeed接口传此字段无效。取值范围如下：1：已开通 2：待开通 3：不允许开通（KA客户）
	UaStat int `json:"uaStat,omitempty"`
	// ValidFlows 可投放流量。此字段仅用于查询，updateAccountFeed接口传此字段无效。数组为空时表示无可投放流量。流量类型与推广单元对象的ftypes字段定义保持一致，取值范围如下：1：手机百度 2：贴吧 4：百青藤 8：好看视频 64：百度小说
	ValidFlows []int `json:"validFlow,omitempty"`
	// TradeId 用户行业ID
	TradeId uint64 `json:"tradeId,omitempty"`
	// BudgetOfflineTime 账户预算撞线时间
	BudgetOfflineTime []map[string]string `json:"budgetOfflineTime,omitempty"`
	// AdType 账户资金包类型
	// 取值范围：枚举值，列表如下
	// 187 - KA客户标识，使用KA原生资金池
	// 188 - 中小客户中的直销/分公司客户
	// 189 - 中小客户中的代理商客户
	AdType int `json:"adType,omitempty"`
}
