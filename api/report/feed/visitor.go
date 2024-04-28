package feed

import (
	reportApi "github.com/xiaoshouchen/baidu-marketing/api/report"
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// Visitor 信息流访客明细报告
func Visitor(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 6759418
	return reportApi.GetReportData(clt, auth, reportRequest)
}
