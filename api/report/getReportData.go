package report

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// GetReportData 一站式多渠道报告
func GetReportData(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	req := &model.Request{
		Header: auth,
		Body:   reportRequest,
	}
	var resp report.GetReportDataResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
