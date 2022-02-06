package report

// GetReportStateResponse 获取异步报告状态 API Response
type GetReportStateResponse struct {
	Data []struct {
		IsGenerated int `json:"isGenerated,omitempty"` // 报告生成状态; 1 - 等待中; 2 - 处理中; 3 - 处理成功
	} `json:"data,omitempty"`
}
