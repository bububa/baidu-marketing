package search

import (
	"context"

	reportApi "github.com/bububa/baidu-marketing/api/report"
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// Audience 人群报告
func Audience(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 9718404
	return reportApi.GetReportData(ctx, clt, auth, reportRequest)
}
