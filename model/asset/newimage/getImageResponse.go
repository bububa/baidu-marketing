package newimage

// GetImageResponse （新图片服务（ImageManageService）打通了搜索图库与信息流图库）获取图片素材 API Response
type GetImageResponse struct {
	// ListData 返回数据对象
	ListData []Image `json:"listData,omitempty"`
	// TotalCount 数据总数
	TotalCount int `json:"totalCount,omitempty"`
	// PageNo 当前页
	PageNo int `json:"pageNo,omitempty"`
}
