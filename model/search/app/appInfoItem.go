package app

type InfoAppItem struct {
	CampaignId     int     `json:"campaignId"`
	AdgroupId      int     `json:"adgroupId"`
	Name           string  `json:"name"`
	Version        string  `json:"version"`
	Platform       int     `json:"platform"`
	CbindId        int     `json:"cbindId"`
	ChannelPackage string  `json:"channelPackage"`
	Status         int     `json:"status"`
	BidRatio       float64 `json:"bidRatio"`
	AppSource      int     `json:"appSource"`
	ChannelId      int64   `json:"channelId"`
	AppStoreId     int64   `json:"appStoreId"`
	PackageName    string  `json:"packageName"`
}
