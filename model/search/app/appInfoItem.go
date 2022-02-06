package app

// AppInfoItem APP对象
type AppInfoItem struct {
	// Name app名称
	Name string `json:"name,omitempty"`
	// Version app版本
	Version string `json:"version,omitempty"`
	// Platform app操作系统
	Platform int `json:"platform,omitempty"`
	// VersionId app版本Id
	VersionId int `json:"versionId,omitempty"`
	// AppId app唯一标识
	AppId int `json:"appId,omitempty"`
	// PackageName app包名
	PackageName string `json:"packageName,omitempty"`
	// Status app状态
	Status int `json:"status,omitempty"`
	// AppSource app来源
	AppSource int `json:"appSource,omitempty"`
	// ChannelId Android包唯一标识
	ChannelId int `json:"channelId,omitempty"`
	// AppStoreId IOS包唯一标识
	AppStoreId int `json:"appStoreId,omitempty"`
	// Icon app图标
	Icon string `json:"icon,omitempty"`
	// DownloadUrl app下载地址
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// ChannelName app渠道包名称
	ChannelName string `json:"channelName,omitempty"`
}
