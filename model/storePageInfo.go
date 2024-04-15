package model

// StorePageInfo 本地计划设置的门店落地页信息
type StorePageInfo struct {
	// StoreId 门店id
	StoreId uint64 `json:"storeId,omitempty"`
	// PageId 落地页id
	PageId uint64 `json:"pageId,omitempty"`
	// Url 落地页url
	Url string `json:"url,omitempty"`
	// PageType 落地页类型 0 梧桐 1 h5+小程序 2 医美 3 旺铺
	PageType int `json:"pageType,omitempty"`
	// MonitorCode 监控代码
	MonitorCode string `json:"monitorCode,omitempty"`
}
