package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetReportFeedStateRequest 获取异步报告状态 API Request
type GetReportFeedStateRequest struct {
	// ReportId 报告ID
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFeedStateRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedState", model.BASE_URL_FEED)
}
