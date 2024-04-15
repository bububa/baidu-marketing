package account

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetAccountInfoRequest 查询账户 API Request
type GetAccountInfoRequest struct {
	// AccountFields 指定需要返回的属性;
	// 取值范围：枚举值，列表如下
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
	// geoLocationStatus - 推广地域地理位置选项
	// excludeQueryRegionStatus - 搜索意图地域词排除
	// longMonitorSublink - 创意组件-文字链监控代码
	// accountMonitorUrl - 账号-通用点击监测地址
	// cid - 账号主体ID
	// liceName - 账号主体名称
	AccountFields []string `json:"accountFields,omitempty"`
}

func (r GetAccountInfoRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AccountService/getAccountInfo")
}

// GetAccountInfoResponse 查询账户 API Response
type GetAccountInfoResponse struct {
	Data []Account `json:"data,omitempty"`
}
