package adgroup

// Adgroup 推广单元
type Adgroup struct {
	// AdgroupId 推广单元ID
	AdgroupId uint64 `json:"adgroupId,omitempty"`
	// CampaignId 计划ID
	CampaignId uint64 `json:"campaignId,omitempty"`
	// AdgroupName 单元名称;最大30个字节（1个中文按2个字节计算，英文、数字按1个字节计算）
	AdgroupName string `json:"adgroupName,omitempty"`
	// MaxPrice 单元出价;取值范围：(0,999.99] &&<= 所属计划预算
	MaxPrice *float64 `json:"maxPrice,omitempty"`
	// Pause 暂停状态;
	//     取值范围：枚举值，列表如下
	// true - 暂停
	// false - 启用
	Pause *bool `json:"pause,omitempty"`
	// NegativeWords 单元短语否定关键词;单个否词最长40字节(1个中文按2个字节计算，英文、数字按1个字节计算)，数组元素个数最多200个
	NegativeWords []string `json:"negativeWords,omitempty"`
	// ExactNegativeWords 单元精确否定关键词;单个否词最长40字节(1个中文按2个字节计算，英文、数字按1个字节计算)，数组元素个数最多200个
	ExactNegativeWords []string `json:"exactNegativeWords,omitempty"`
	// Status 单元状态
	// 取值范围：枚举值，列表如下
	// 31 - 有效
	// 32 - 暂停推广
	// 33 - 推广计划暂停推广
	// 43 - 未审核
	Status int `json:"status,omitempty"`
	// AdType 广告类型
	// 取值范围：枚举值，列表如下
	// 0 - 普通单元
	// 14 - 商品单元
	AdType *int `json:"adType,omitempty"`
	// AppShopDirectStatus 应用商店直投
	// 仅支持应用推广营销目标
	// 默认值：关闭
	// 取值范围：枚举值，列表如下
	// 0 - 关闭
	// 1 - 开启
	AppShopDirectStatus *int `json:"appShopDirectStatus,omitempty"`
	// SegmentRecommendStatus 基础创意智能配图开关
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 开启
	// 1 - 关闭
	SegmentRecommendStatus *int `json:"segmentRecommendStatus,omitempty"`
	// ProductSetId 虚拟商品组id;计划类型为商品计划时必填，使用DpaProductSetService服务创建商品组
	ProductSetId int64 `json:"productSetId,omitempty"`
	// PaPrice 推广单元商品出价;计划类型为商品计划时必填，优先级高于maxPrice。商品组中每个商品每次展现并被点击的最高费用取值范围：(0,999.99]
	PaPrice *float64 `json:"paPrice,omitempty"`
	// MonitorUrl 单元层级监控url;仅计划类型为商品计划时支持。单元层级监控url，附加到所有样式中作为监控url。
	MonitorUrl *string `json:"monitorUrl,omitempty"`
	// PcFinalUrl 计算机最终访问网址
	PcFinalUrl *string `json:"pcFinalUrl,omitempty"`
	// PcTrackParam 计算机监控后缀
	PcTrackParam *string `json:"pcTrackParam,omitempty"`
	// PcTrackTemplate 计算机第三方追踪模板
	PcTrackTemplate *string `json:"pcTrackTemplate,omitempty"`
	// MobileFinalUrl 移动最终访问网址
	MobileFinalUrl *string `json:"mobileFinalUrl,omitempty"`
	// MobileTrackParam 移动监控后缀
	MobileTrackParam *string `json:"mobileTrackParam,omitempty"`
	// MobileTrackTemplate 移动第三方追踪模板
	MobileTrackTemplate *string `json:"mobileTrackTemplate,omitempty"`
	// AdgroupAutoTargetingStatus 自动定向
	// 默认值：true
	// 取值范围：枚举值，列表如下
	// true - 开启
	// false - 关闭
	AdgroupAutoTargetingStatus *bool `json:"adgroupAutoTargetingStatus,omitempty"`
	// CreativeTextOptimizationStatus 自动文案优化
	// 默认值：true
	// 取值范围：枚举值，列表如下
	// true - 开启
	// false - 关闭
	CreativeTextOptimizationStatus *bool `json:"creativeTextOptimizationStatus,omitempty"`
	// GetTemp 是否查询单元影子
	// 默认值：0
	// 0 - 只查询单元本身，1 - 只查询单元影子；
	// 想要获得单元的全集，需要调用该方法两次，分别为getTemp=0和getTemp=1；
	// 影子说明：用户先向系统提交了单元A，并且A已审核通过，之后再对A进行影响审核状态的修改（修改pcFinalUrl、pcTrackParam、pcTrackTemplate、mobileFinalUrl、mobileTrackParam、mobileTrackTemplate等字段），修改后的单元为A’（A’即为影子，仅对审核通过的物料进行修改才会产生影子），在A’通过审核生效之前，线上的生效创意仍然为A。 此时：getTemp为0查询到的是A，getTemp为1查询到的是A’
	GetTemp *bool `json:"getTemp,omitempty"`
	// AdgroupAppBinds 单元与app的绑定信息
	AdgroupAppBinds []AdgroupAppBindInfo `json:"adgroupAppBinds,omitempty"`
}

// AdgroupAppBindInfo 单元与app的绑定信息
type AdgroupAppBindInfo struct {
	// AndroidBindType Android App绑定信息
	AndroidBindType *AndroidAppBindType `json:"androidBindType,omitempty"`
	// IosBindType IOS App绑定信息
	IosBindType *IosAppBindType `json:"iosBindType,omitempty"`
}

// AndroidAppBindType Android App绑定信息
type AndroidAppBindType struct {
	// Platform 设备类型
	Platform int `json:"platform,omitempty"`
	// ChannelId Android的APP包Id, Android的APP唯一标识
	ChannelId uint64 `json:"channelId,omitempty"`
}

// IosAppBindType IOS App绑定信息
type IosAppBindType struct {
	// Platform 设备类型
	Platform int `json:"platform,omitempty"`
	// AppStoreId IOS包唯一标识
	AppStoreId uint64 `json:"appStoreId,omitempty"`
}
