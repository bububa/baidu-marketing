package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetReportFileUrlRequest struct {
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFileUrlRequest) Url() string {
	return fmt.Sprintf("%sReportService/getReportFileUrl", model.BASE_URL_SMS)
}
