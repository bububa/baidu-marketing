package feed

import (
	"context"

	reportApi "github.com/bububa/baidu-marketing/api/report"
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// Video 视频报告
func Video(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 114718
	return reportApi.GetReportData(ctx, clt, auth, reportRequest)
}
