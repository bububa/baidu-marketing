package app

// GetAppListResponse 获取APP素材 API Response
type GetAppListResponse struct {
	Data []AppInfoItemList `json:"data,omitempty"`
}

type AppInfoItemList struct {
	AppInfoList []AppInfoItem `json:"appInfoList,omitempty"`
	TotalCount  int           `json:"totalCount,omitempty"`
}
