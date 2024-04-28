package search

import (
	reportApi "github.com/xiaoshouchen/baidu-marketing/api/report"
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// QueryWord 搜索词报告
func QueryWord(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2307838
	return reportApi.GetReportData(clt, auth, reportRequest)
}
