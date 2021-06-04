package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetProfessionalReportIdRequest struct {
	ReportRequestType *ReportRequest `json:"reportRequestType,omitempty"`
}

func (r GetProfessionalReportIdRequest) Url() string {
	return fmt.Sprintf("%sReportService/getProfessionalReportId", model.BASE_URL_SMS)
}
