package account

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetAccountFeedRequest 查询账户信息 API Request
type GetAccountFeedRequest struct {
	// AccountFeedFields 需要查询的账户属性。取值范围参考账户信息对象中的字段取值
	// 取值范围：枚举值，列表如下
	// userId - 账户ID
	// balance - 账户余额
	// budget - 账户预算
	// balancePackage - 资金包类型
	// userStat - 账户状态
	// uaStatus - 是否开通feed产品线权限
	// validFlows - 可投放流量
	// cid - 账户主体ID
	// liceName - 账户主体名称
	// tradeId - 用户行业ID
	// budgetOfflineTime - 账户预算撞线时间
	// adtype - 账户资金包类型
	AccountFeedFields []string `json:"accountFeedFields,omitempty"`
}

func (r GetAccountFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AccountFeedService/getAccountFeed")
}

// GetAccountFeedResponse 查询账户信息 API Response
type GetAccountFeedResponse struct {
	Data []Account `json:"data,omitempty"`
}
