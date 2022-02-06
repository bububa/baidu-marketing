package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/report"
)

// GetReportFeedFileUrl 获取异步报告文件URL
// 获取报告下载地址。当报告成功生成后，使用reportId请求，返回相应的报告下载地址。
func GetReportFeedFileUrl(clt *core.SDKClient, auth model.RequestHeader, reportId string) (*model.ResponseHeader, string, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetReportFeedFileUrlRequest{
			ReportId: reportId,
		},
	}
	var resp report.GetReportFeedFileUrlResponse
	header, err := clt.Do(req, &resp)
	if err != nil && header != nil && header.Status != 1 {
		return header, "", err
	}
	var retPath string
	if len(resp.Data) > 0 {
		retPath = resp.Data[0].ReportFilePath
	}
	return header, retPath, err
}
