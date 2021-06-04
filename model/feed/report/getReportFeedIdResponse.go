package report

type GetReportFeedIdResponse struct {
	Data struct {
		ReportId string `json:"reportId,omitempty"`
	} `json:"data,omitempty"`
}
