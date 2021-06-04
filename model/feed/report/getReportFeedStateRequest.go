package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetReportFeedStateRequest struct {
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFeedStateRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedState", model.BASE_URL_FEED)
}
