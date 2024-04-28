package feed

import (
	reportApi "github.com/xiaoshouchen/baidu-marketing/api/report"
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// Overall 信息流整体账户报告
func Overall(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2172649
	return reportApi.GetReportData(clt, auth, reportRequest)
}
