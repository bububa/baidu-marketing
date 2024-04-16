package app

// AppBindItem APP绑定对象
type AppBindItem struct {
	// CampaignId 推广计划id
	CampaignId uint64 `json:"campaignId,omitempty"`
	// AdgroupId 推广单元id
	AdgroupId uint64 `json:"adgroupId,omitempty"`
	// Name app名称
	Name string `json:"name,omitempty"`
	// Version app版本
	Version string `json:"version,omitempty"`
	// Platform app操作系统
	// 取值范围：枚举值，列表如下
	// 1 - 安卓
	// 3 - iOS
	Platform int `json:"platform,omitempty"`
	// CbindId 绑定关系id
	CbindId uint64 `json:"cbindId,omitempty"`
	// ChannelPackage app渠道包名称
	ChannelPackage string `json:"channelPackage,omitempty"`
	// Status 状态
	// 第一位表示暂停位（0启动，1暂停）；第二位表示APP有效位（0有效，1无效）；第三位表示审核中（0非审核中，1审核中）
	Status int `json:"status,omitempty"`
	// AppSource app来源
	// 取值范围：枚举值，列表如下
	// 0 - 百度移动应用平台
	// 1 - 极速下载
	// 2 - 投放平台
	// 3 - 百度移动应用平台
	AppSource int `json:"appSource,omitempty"`
	// ChannelId Android包唯一标识
	ChannelId uint64 `json:"channelId,omitempty"`
	// AppStoreId IOS包唯一标识
	AppStoreId uint64 `json:"appStoreId,omitempty"`
	// PackageName app包名
	PackageName string `json:"packageName,omitempty"`
}
