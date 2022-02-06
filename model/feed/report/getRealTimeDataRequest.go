package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetRealTimeDataRequest 实时报告请求 API Request
type GetRealTimeDataRequest struct {
	RealTimeRequestType *RealTimeRequest `json:"realTimeRequestType"`
}

func (r GetRealTimeDataRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getRealTimeFeedData", model.BASE_URL_FEED)
}
