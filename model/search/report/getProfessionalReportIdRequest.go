package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetProfessionalReportIdRequest 创建异步报告（获取异步报告id） API Request
type GetProfessionalReportIdRequest struct {
	ReportRequestType *ReportRequest `json:"reportRequestType,omitempty"`
}

func (r GetProfessionalReportIdRequest) Url() string {
	return fmt.Sprintf("%sReportService/getProfessionalReportId", model.BASE_URL_SMS)
}
