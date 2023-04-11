package campaign

// App 推广app信息
type App struct {
	// AppName 应用名称，最多20个字符。
	AppName string `json:"appName,omitempty"`
	// ApkName 应用包名，仅Android有效，最多1024个字符
	ApkName string `json:"apkName,omitempty"`
	// AppUrl 推广应用链接。仅支持IOS app，链接必须为iTunes链接。
	AppUrl string `json:"appUrl,omitempty"`
	// DocId 仅支持Android app
	DocId int64 `json:"docId,omitempty"`
	// ChannelId 仅支持Android app, 请登录移动开发者平台查看推广APP对应的channelId信息
	ChannelId int64 `json:"channelId,omitempty"`
	// OpenUrl 应用商店直投链接，选填，须以market://details?或apps://ProductDetail/开头；或者"0"表示关闭,“1”表示开启. 更新时不填表示清除已有链接
	OpenUrl string `json:"openUrl,omitempty"`
	// DownloadType 下载方式。0: 直接下载 1：落地页下载
	DownloadType int `json:"downloadType,omitempty"`
}
