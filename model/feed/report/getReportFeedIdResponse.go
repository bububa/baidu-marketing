package report

// GetReportFeedIdResponse 创建异步报告（获取异步报告id） API Response
type GetReportFeedIdResponse struct {
	Data struct {
		// ReportId 异步报告ID。比如：8e7e3f2d84a19c5df1415957434b2bd8; 可用此ID调用查询异步报告状态和文件URL。
		ReportId string `json:"reportId,omitempty"`
	} `json:"data,omitempty"`
}
