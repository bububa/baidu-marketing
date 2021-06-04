package report

import (
	"errors"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/report"
)

// 获取异步报告状态
// 查询报告当前的生成状态。请求中提供报告ID，返回报告的处理状态。
// 说明：在获取Report文件url前，请调用此方法。待确认报表已生成时，再获取下载的url
func GetReportFeedState(clt *core.SDKClient, auth model.RequestHeader, reportId string) (int, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetReportFeedStateRequest{
			ReportId: reportId,
		},
	}
	var resp report.GetReportFeedStateResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return 0, err
	}
	if len(resp.Data) == 0 {
		return 0, errors.New("no result")
	}
	return resp.Data[0].IsGenerated, nil
}
