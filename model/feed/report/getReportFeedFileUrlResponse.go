package report

// GetReportFeedFileUrlResponse 获取异步报告文件url API Response
type GetReportFeedFileUrlResponse struct {
	Data []struct {
		// ReportFilePath 报告问价下载地址，生成的url有效期为1小时，如果超时则需重新获取新的url。
		ReportFilePath string `json:"reportFilePath,omitempty"`
	} `json:"data,omitempty"`
}
