package adgroup

/* Adgroup 推广单元
// Pause 可选值
// true - 暂停
// false - 启用

// Status 可选值
// 31 - 有效
// 32 - 暂停推广
// 33 - 推广计划暂停推广

// PriceRatio 可选值
// 计划为计算机优先（bidprefer=1）时有效，取值范围：[0-10]，默认为1；
// bidprefer=2时仅能为默认值null；
// 查询时返回-1表示未设置；

// PcPriceRatio 可选值
// 计划为移动优先（bidprefer=2）时有效，取值范围：[0-10]，默认值为1；
// bidprefer=1时仅能为默认值null；
// 查询时返回-1表示未设置

// AdType 可选值
// 0 - 普通单元
// 14 - 商品单元
// 6 - 网址定向单元

// SegmentRecommendStatus 可选值
// 0 - 开启
// 1 - 关闭
*/
type Adgroup struct {
	// AdgroupId 推广单元ID
	AdgroupId int64 `json:"adgroupId,omitempty"`
	// CampaignId 计划ID
	CampaignId int64 `json:"campaignId,omitempty"`
	// AdgroupName 单元名称;最大30个字节（1个中文按2个字节计算，英文、数字按1个字节计算）
	AdgroupName string `json:"adgroupName,omitempty"`
	// MaxPrice 单元出价;取值范围：(0,999.99] &&<= 所属计划预算
	MaxPrice float64 `json:"maxPrice,omitempty"`
	// Pause 暂停状态;
	Pause *bool `json:"pause,omitempty"`
	// NegativeWords 单元短语否定关键词;单个否词最长40字节(1个中文按2个字节计算，英文、数字按1个字节计算)，数组元素个数最多200个
	NegativeWords []string `json:"negativeWords,omitempty"`
	// ExactNegativeWords 单元精确否定关键词;单个否词最长40字节(1个中文按2个字节计算，英文、数字按1个字节计算)，数组元素个数最多200个
	ExactNegativeWords []string `json:"exactNegativeWords,omitempty"`
	// Status 单元状态
	Status int `json:"status,omitempty"`
	// PriceRatio 单元移动出价系数（默认使用计划的）
	PriceRatio float64 `json:"priceRatio,omitempty"`
	// PcPriceRatio 单元计算机出价系数（默认使用计划的）
	PcPriceRatio float64 `json:"pcPriceRatio,omitempty"`
	// AdType 广告类型
	AdType *int `json:"adType,omitempty"`
	// SegmentRecommendStatus 基础创意智能配图开关
	SegmentRecommendStatus *int `json:"segmentRecommendStatus,omitempty"`
	// ProductSetId 虚拟商品组id;计划类型为商品计划时必填，使用DpaProductSetService服务创建商品组
	ProductSetId int64 `json:"productSetId,omitempty"`
	// PaPrice 推广单元商品出价;计划类型为商品计划时必填，优先级高于maxPrice。商品组中每个商品每次展现并被点击的最高费用取值范围：(0,999.99]
	PaPrice float64 `json:"paPrice,omitempty"`
	// MonitorUrl 单元层级监控url;仅计划类型为商品计划时支持。单元层级监控url，附加到所有样式中作为监控url。
	MonitorUrl string `json:"monitorUrl,omitempty"`
}
