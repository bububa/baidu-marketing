package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetReportFeedIdRequest 创建异步报告（获取异步报告id）
type GetReportFeedIdRequest struct {
	// ReportRequestType 报告查询条件
	ReportRequestType *ReportRequest `json:"reportRequestType,omitempty"`
}

func (r GetReportFeedIdRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getReportFeedId", model.BASE_URL_FEED)
}
