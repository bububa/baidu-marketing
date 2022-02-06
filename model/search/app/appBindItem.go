package app

// AppBindItem APP绑定对象
type AppBindItem struct {
	// CampaignId 推广计划id
	CampaignId int `json:"campaignId,omitempty"`
	// AdgroupId 推广单元id
	AdgroupId int `json:"adgroupId,omitempty"`
	// Name app名称
	Name string `json:"name,omitempty"`
	// Version app版本
	Version string `json:"version,omitempty"`
	// Platform app操作系统
	Platform int `json:"platform,omitempty"`
	// CbindId 绑定关系id
	CbindId int `json:"cbindId,omitempty"`
	// ChannelPackage app渠道包名称
	ChannelPackage string `json:"channelPackage,omitempty"`
	// Status 状态
	Status int `json:"status,omitempty"`
	// BidRatio 动态溢价系数
	BidRatio float64 `json:"bidRatio,omitempty"`
	// AppSource app来源
	AppSource int `json:"appSource,omitempty"`
	// ChannelId Android包唯一标识
	ChannelId int64 `json:"channelId,omitempty"`
	// AppStoreId IOS包唯一标识
	AppStoreId int64 `json:"appStoreId,omitempty"`
	// PackageName app包名
	PackageName string `json:"packageName,omitempty"`
}
