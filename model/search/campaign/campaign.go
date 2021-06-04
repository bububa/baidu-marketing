package campaign

import "github.com/bububa/baidu-marketing/model"

// Status 可选值
// 21 - 有效
// 22 - 处于暂停时段
// 23 - 暂停推广
// 24 - 计划预算不足
// 25 - 账户预算不足
//
// PriceRatio 可选值
// bidPrefer=1时有效，取值范围：[0, 10]，支持精确到两位小数，默认值为1；
// bidPrefer=2时，无效字段，getCampaign请求返回默认值1
//
// BidPrefer 可选值
// 默认值：1
// 取值范围：枚举值，列表如下
// 1 - 以计算机出价为基准
// 2 - 以移动出价为基准
//
// AdType 可选值
// 默认值：0
// 取值范围：枚举值，列表如下
// 0 - 普通计划
// 14 - 商品计划
// 6 - 网址定向计划
//
// PaDevice 可选值
// 默认值：0
// 取值范围：枚举值，列表如下
// 0 - 全部
// 1 - 计算机
// 2 - 移动
//
// OS 可选值
// 默认值：全选
// 取值范围：枚举值，列表如下
// IPHONE - 苹果手机
// ANDROID - 安卓手机
// OTHERS - 其他类型
//
// MarketingTargetId 可选值
// 0 - 网站链接
// 1 - 应用推广
// 2 - 门店推广
// 4 - 电商店铺推广
// 5 - 商品目录
//
// ShopType 可选值
// 1 - 度小店
// 3 - 第三方店铺
// 31 - 淘宝（含天猫）
// 32 - 京东
// 33 - 拼多多
// 34 - 苏宁
// 当营销目标为"电商店铺推广"时必填，其他营销目标不支持
type Campaign struct {
	CampaignId            int64                       `json:"campaignId,omitempty"`            // 计划ID
	CampaignName          string                      `json:"campaignName,omitempty"`          // 计划名称;长度限制：最大30个字节（1个中文按2个字节计算，英文、数字按1个字节计算）
	Budget                float64                     `json:"budget,omitempty"`                // 计划每日预算;取值范围：[50, Min(10000000, 账户预算)]
	RegionTarget          []int                       `json:"regionTarget,omitempty"`          // 计划推广地域
	NegativeKeywords      []string                    `json:"negativeKeywords,omitempty"`      // 短语否定关键词列表
	ExactNegativeKeywords []string                    `json:"exactNegativeKeywords,omitempty"` // 精确否定关键词列表
	Schedule              []model.Schedule            `json:"schedule,omitempty"`              // 计划推广暂停时段
	BudgetOfflineTime     []model.OfflineTime         `json:"budgetOfflineTime,omitempty"`     // 预算下线时间;数组元素个数限制：最近有过下线时段的7个自然日的下线和上线时段（这7个自然日中若某日期距当前已超过30天，则不返回）;
	ShowProb              int                         `json:"showProb,omitempty"`              // 创意展现方式;1 - 优选;2 - 轮替
	Pause                 bool                        `json:"pause,omitempty"`                 // 暂停状态;true - 暂停;false - 启用
	Status                int                         `json:"status,omitempty"`                // 计划状态
	PriceRatio            float64                     `json:"priceRatio,omitempty"`            // 移动出价系数
	PcPriceRatio          float64                     `json:"pcPriceRatio,omitempty"`          // 计算机出价系数
	BidPrefer             int                         `json:"bidPrefer,omitempty"`             // 计划出价类型
	AdType                int                         `json:"adType,omitempty"`                // 计划类型
	BusinessPointId       int64                       `json:"businessPointId,omitempty"`       // 推广业务ID
	BusinessPointName     string                      `json:"businessPointName,omitempty"`     // 推广业务字面
	SmartRegion           bool                        `json:"smartRegion,omitempty"`           // 商品计划: 智能地域开关
	PaDevice              int                         `json:"paDevice,omitempty"`              // 商品计划: 计划的投放设备
	Os                    []string                    `json:"os,omitempty"`                    // 商品计划: 计划的投放设备平台
	RegionPriceFactor     []model.RegionPriceFactor   `json:"regionPriceFactor,omitempty"`     // 分地域出价系数
	SchedulePriceFactor   []model.SchedulePriceFactor `json:"schedulePriceFactor,omitempty"`   // 分时段出价系数
	MarketingTargetId     int                         `json:"marketingTargetId,omitempty"`     // 营销目标类型
	ShopType              int                         `json:"shopType,omitempty"`              // 电商店铺类型
}
