package account

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateAccountInfoRequest 更新账户 API Request
type UpdateAccountInfoRequest struct {
	// AccountInfo 更新账户信息
	AccountInfo *UpdateAccountInfo `json:"accountInfo,omitempty"`
}

func (r UpdateAccountInfoRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AccountService/updateAccountInfo")
}

// UpdateAccountInfoResponse 更新账户信息 API Response
type UpdateAccountInfoResponse struct {
	Data []Account `json:"data,omitempty"`
}

// UpdateAccountInfo 更新账户信息
type UpdateAccountInfo struct {
	// Budget 账户预算, 当设置为日预算时，取值范围：[50, 10000000]； 当不设置预算时，输入任意值均默认为0
	Budget *float64 `json:"budget,omitempty"`
	// RegionTarget 推广地域列表
	RegionTarget []int64 `json:"regionTarget,omitempty"`
	// ExcludeIps IP排除列表, 支持添加/删除ipv4和ipv6类型ip, 例如：2001:0DB8:0000:0023:0008:0800:200C:417A
	ExcludeIps []string `json:"excludeIps,omitempty"`
	// BudgetType 预算类型; 0 - 不设置预算;1 - 日预算
	BudgetType *int `json:"budgetType,omitempty"`
	// RegionPriceFactor 分地域出价系数
	RegionPriceFactor []model.RegionPriceFactor `json:"regionPriceFactor,omitempty"`
	// GeolocationStatus 推广地域地理位置选项 取值范围：枚举值，列表如下
	// 0 - 该地区内或搜索意图在该地区的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户、在搜索词中对该地区表现出明确兴趣的用户）
	// 1 - 该地区内的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户）
	// 原账户层级的“搜索意图地域词扩展”已升级为计划地域设置中的推广地域地理位置选项
	GeolocationStatus *int `json:"geoLocationStatus,omitempty"`
	// ExcludeQueryRegionStatus 搜索意图地域词排除; 取值范围：枚举值，列表如下
	// true - 启用
	// false - 关闭
	ExcludeQueryRegionStatus *bool `json:"excludeQueryRegionStatus,omitempty"`
	// LongMonitorSublink 创意组件-文字链监控代码; 可填充监控代码默认值："product=xinxizhenlie_longsublink"，如需自定义，请直接录入监控代码（字符限制255）
	LongMonitorSublink *string `json:"longMonitorSublink,omitempty"`
	// AccountMonitorURL 通用点击监测地址; 当前通用点击检测地址长度限制1024
	AccountMonitorURL *string `json:"accountMonitorURL,omitempty"`
}
