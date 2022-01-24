package app

type InfoApp struct {
	AppInfoList []struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Platform    int    `json:"platform"`
		VersionId   int    `json:"versionId"`
		AppId       int    `json:"appId"`
		PackageName string `json:"packageName"`
		Status      int    `json:"status"`
		AppSource   int    `json:"appSource"`
		ChannelId   int    `json:"channelId"`
		AppStoreId  int    `json:"appStoreId"`
		Icon        string `json:"icon"`
		DownloadUrl string `json:"downloadUrl"`
		ChannelName string `json:"channelName"`
	} `json:"appInfoList"`
	TotalCount int `json:"totalCount"`
}
