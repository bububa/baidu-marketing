package report

// GetProfessionalReportIdResponse 创建异步报告（获取异步报告id） API Response
type GetProfessionalReportIdResponse struct {
	Data struct {
		ReportId string `json:"reportId,omitempty"`
	} `json:"data,omitempty"`
}
