package report

// GetReportFeedStateResponse 获取异步报告状态
type GetReportFeedStateResponse struct {
	Data []struct {
		// IsGenerated 报告生成状态; 1 - 等待中; 2 - 处理中; 3 - 处理成功
		IsGenerated int `json:"isGenerated,omitempty"`
	} `json:"data,omitempty"`
}
