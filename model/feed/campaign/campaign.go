package campaign

import "github.com/bububa/baidu-marketing/model"

// Campaign 计划对象
/*
推广计划状态。此字段仅用于查询接口响应结果。取值范围：枚举值，列表如下
0 - 有效
1 - 处于暂停时段
2 - 暂停推广
3 - 推广计划预算不足
4 - 账户待激活
11 - 账户预算不足
20 - 账户余额为零
23 - 被禁推
24 - app已下线
25 - 应用审核中
*/
type Campaign struct {
	// CampaignId 计划ID
	CampaignId int64 `json:"campaignFeedId,omitempty"`
	// CampaignName 计划名称。长度限制最大100个字节，1个中文及中文符号按2个字节计算
	CampaignName string `json:"campaignFeedName,omitempty"`
	// Subject 推广对象。取值范围如下：1：网站链接2：应用下载（IOS）3：应用下载（Android）
	Subject int `json:"subject,omitempty"`
	// AppInfo 推广app信息。subject=1时，该字段无效。对象定义参考下文推广app信息
	AppInfo *App `json:"appInfo,omitempty"`
	// Budget 推广计划预算。默认为0,表示不限预算。正常取值范围为[50 - 9999999.99]
	Budget float64 `json:"budget,omitempty"`
	// StartTime 推广开始日期。默认为null，表示长期投放。格式示例：'2016-12-15'不能早于当天的日期
	StartTime string `json:"starttime,omitempty"`
	// EndTime 推广结束日期。默认为null，表示长期投放。例如：'2016-12-18'不能早于开始日期
	EndTime string `json:"endtime,omitempty"`
	// Schedule 暂停时段设置，对象定义参考下文暂停时段设置
	Schedule []model.Schedule `json:"schedule,omitempty"`
	// BgtctlType 预算分配控制方式。默认为标准方式取值范围如下：0：匀速。根据流量波动，让预算在整个投放日程中较为平稳的消耗。1：标准。尽快将广告投放出去，预算可能会在短时间内消耗完2：加速。尽可能获得更多展现，对比标准投放预算消耗更快
	BgtctlType int `json:"bgtctltype,omitempty"`
	// Pause 是否暂停推广。默认为false。true：推广计划暂停 false：推广计划启用
	Pause *bool `json:"pause,omitempty"`
	// Status 推广计划状态。
	Status int `json:"status,omitempty"`
	// BsType 1：普通计划 3：闪投计划 7：原生RTA 注：不支持修改
	BsType int `json:"bstype,omitempty"`
	// CampaignType 信息流计划类型。1： 普通模式 4：放量模式
	CampaignType int `json:"campaignType,omitempty"`
	// EshopType 交易所在平台。取值范围：枚举值，列表如下 3-1 淘宝（含天猫） 3-2 京东 3-3 拼多多 3-4 苏宁 仅推广对象为电商店铺时需传该字段
	EshopType string `json:"eshopType,omitempty"`
	// AddTime 添加时间
	AddTime string `json:"addtime,omitempty"`
}
