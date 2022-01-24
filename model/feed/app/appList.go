package app

type ListApp struct {
	AppName        string      `json:"appName"`
	ApkName        string      `json:"apkName"`
	AppUrl         string      `json:"appUrl"`
	DocId          int         `json:"docId"`
	ChannelId      int         `json:"channelId"`
	ChannelPackage interface{} `json:"channelPackage"`
	AppStatus      int         `json:"appStatus"`
}
