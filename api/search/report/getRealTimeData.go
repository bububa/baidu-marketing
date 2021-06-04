package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// 推广报告
func GetRealTimeData(clt *core.SDKClient, auth model.RequestHeader, realTimeRequest *report.RealTimeRequest) ([]report.RealTimeResult, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetRealTimeDataRequest{
			RealTimeRequestType: realTimeRequest,
		},
	}
	var resp report.GetRealTimeDataResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
