package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetReportFileUrlRequest 获取异步报告文件URL API Request
type GetReportFileUrlRequest struct {
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFileUrlRequest) Url() string {
	return fmt.Sprintf("%sReportService/getReportFileUrl", model.BASE_URL_SMS)
}
