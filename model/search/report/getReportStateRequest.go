package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetReportStateRequest 获取异步报告状态 API Request
type GetReportStateRequest struct {
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportStateRequest) Url() string {
	return fmt.Sprintf("%sReportService/getReportState", model.BASE_URL_SMS)
}
