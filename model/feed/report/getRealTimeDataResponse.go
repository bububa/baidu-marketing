package report

// GetRealTimeDataResponse 实时报告请求 API Response
type GetRealTimeDataResponse struct {
	Data []RealTimeResult `json:"data,omitempty"`
}
