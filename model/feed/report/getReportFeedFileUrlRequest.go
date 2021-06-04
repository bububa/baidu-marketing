package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetReportFeedFileUrlRequest struct {
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFeedFileUrlRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedFileUrl", model.BASE_URL_FEED)
}
