package report

// GetReportFileUrlResponse 获取异步报告文件URL API Response
type GetReportFileUrlResponse struct {
	Data []struct {
		ReportFilePath string `json:"reportFilePath,omitempty"` // 报告问价下载地址，生成的url有效期为1小时，如果超时则需重新获取新的url。
	} `json:"data,omitempty"`
}
