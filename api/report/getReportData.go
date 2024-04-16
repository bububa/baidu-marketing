package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
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
