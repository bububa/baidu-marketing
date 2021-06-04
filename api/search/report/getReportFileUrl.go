package report

import (
	"errors"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// 获取异步报告文件URL
// 获取报告下载地址。当报告成功生成后，使用reportId请求，返回相应的报告下载地址。
func GetReportFileUrl(clt *core.SDKClient, auth model.RequestHeader, reportId string) (string, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetReportFileUrlRequest{
			ReportId: reportId,
		},
	}
	var resp report.GetReportFileUrlResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return "", err
	}
	if len(resp.Data) == 0 {
		return "", errors.New("no result")
	}
	return resp.Data[0].ReportFilePath, nil
}
