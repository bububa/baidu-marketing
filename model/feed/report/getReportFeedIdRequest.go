package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetReportFeedIdRequest struct {
	ReportRequestType *ReportRequest `json:"reportRequestType,omitempty"`
}

func (r GetReportFeedIdRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedId", model.BASE_URL_FEED)
}
