package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// GetReportFileUrl 获取异步报告文件URL
// 获取报告下载地址。当报告成功生成后，使用reportId请求，返回相应的报告下载地址。
func GetReportFileUrl(clt *core.SDKClient, auth model.RequestHeader, reportId string) (*model.ResponseHeader, string, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetReportFileUrlRequest{
			ReportId: reportId,
		},
	}
	var resp report.GetReportFileUrlResponse
	header, err := clt.Do(req, &resp)
	if err != nil && header != nil && header.Status != 1 {
		return header, "", err
	}
	var ret string
	if len(resp.Data) > 0 {
		ret = resp.Data[0].ReportFilePath
	}
	return header, ret, err
}
