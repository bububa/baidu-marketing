package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/report"
)

// GetReportFeedState 获取异步报告状态
// 查询报告当前的生成状态。请求中提供报告ID，返回报告的处理状态。
// 说明：在获取Report文件url前，请调用此方法。待确认报表已生成时，再获取下载的url
func GetReportFeedState(clt *core.SDKClient, auth model.RequestHeader, reportId string) (*model.ResponseHeader, int, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetReportFeedStateRequest{
			ReportId: reportId,
		},
	}
	var resp report.GetReportFeedStateResponse
	header, err := clt.Do(req, &resp)
	if err != nil && header != nil && header.Status != 1 {
		return header, 0, err
	}
	var ret int
	if len(resp.Data) > 0 {
		ret = resp.Data[0].IsGenerated
	}
	return header, ret, err
}
