package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetRealTimeDataRequest struct {
	RealTimeRequestType *RealTimeRequest `json:"realtimeRequestType"`
}

func (r GetRealTimeDataRequest) Url() string {
	return fmt.Sprintf("%sReportService/getRealTimeData", model.BASE_URL_SMS)
}
