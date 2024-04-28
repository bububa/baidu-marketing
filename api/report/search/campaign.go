package search

import (
	reportApi "github.com/xiaoshouchen/baidu-marketing/api/report"
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// Campaign 计划报告
func Campaign(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2290316
	return reportApi.GetReportData(clt, auth, reportRequest)
}
