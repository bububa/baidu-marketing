package feed

import (
	reportApi "github.com/bububa/baidu-marketing/api/report"
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// Video 视频报告
func Video(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 114718
	return reportApi.GetReportData(clt, auth, reportRequest)
}
