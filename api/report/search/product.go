package search

import (
	reportApi "github.com/xiaoshouchen/baidu-marketing/api/report"
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// Product 搜索商品报告
func Product(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 248654
	return reportApi.GetReportData(clt, auth, reportRequest)
}
