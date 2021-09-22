package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetRealTimeDataRequest struct {
	RealTimeRequestType *RealTimeRequest `json:"realTimeRequestType"`
}

func (r GetRealTimeDataRequest) Url() string {
	return fmt.Sprintf("%sReportFeedService/getRealTimeFeedData", model.BASE_URL_FEED)
}
