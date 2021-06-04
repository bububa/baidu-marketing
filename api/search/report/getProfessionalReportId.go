package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// 创建异步报告（获取异步报告id）
// 创建异步报告，获取报告ID(reportId)
// 基于reportType可创建多种数据报告的异步任务，具体报告规则可参数基础数据报告介绍
func GetProfessionalReportId(clt *core.SDKClient, auth model.RequestHeader, reqBody *report.ReportRequest) (string, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetProfessionalReportIdRequest{
			ReportRequestType: reqBody,
		},
	}
	var resp report.GetProfessionalReportIdResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return "", err
	}
	return resp.Data.ReportId, nil
}
