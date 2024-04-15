package app

// AppInfoItem APP对象
type AppInfoItem struct {
	// Name app名称
	Name string `json:"name,omitempty"`
	// Version app版本
	Version string `json:"version,omitempty"`
	// Platform app操作系统
	// 取值范围：枚举值，列表如下
	// 1 - 安卓
	// 3 - iOS
	Platform int `json:"platform,omitempty"`
	// VersionId app版本Id
	VersionId uint64 `json:"versionId,omitempty"`
	// AppId app唯一标识
	AppId uint64 `json:"appId,omitempty"`
	// PackageName app包名
	PackageName string `json:"packageName,omitempty"`
	// Status app状态
	// 取值范围：枚举值，列表如下
	// 0 - 有效
	// 1 - 无效
	// 2 - 审核中
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
	// Icon app图标
	Icon string `json:"icon,omitempty"`
	// DownloadUrl app下载地址
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// ChannelName app渠道包名称
	ChannelName string `json:"channelName,omitempty"`
}
