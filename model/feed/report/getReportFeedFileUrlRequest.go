package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetReportFeedFileUrlRequest 获取异步报告文件url API Request
type GetReportFeedFileUrlRequest struct {
	// ReportId 报告ID
	ReportId string `json:"reportId,omitempty"`
}

func (r GetReportFeedFileUrlRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedFileUrl", model.BASE_URL_FEED)
}
