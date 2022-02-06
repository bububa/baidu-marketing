package app

// ListApp APP信息
type ListApp struct {
	// AppName 推广APP名称
	AppName string `json:"appName"`
	// ApkName 推广APP包名
	ApkName string `json:"apkName"`
	// AppUrl 推广APP链接
	AppUrl string `json:"appUrl"`
	// DocId 推广APP docId
	DocId int `json:"docId"`
	// ChannelId 渠道包ID
	ChannelId int `json:"channelId"`
	// ChannelPackage 渠道包名称
	ChannelPackage string `json:"channelPackage"`
	// AppStatus 应用审核状态
	AppStatus int `json:"appStatus"`
}
