package report

type GetReportStateResponse struct {
	Data []struct {
		IsGenerated int `json:"isGenerated,omitempty"` // 报告生成状态; 1 - 等待中; 2 - 处理中; 3 - 处理成功
	} `json:"data,omitempty"`
}
