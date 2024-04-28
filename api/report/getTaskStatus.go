package report

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/report"
)

// GetTaskStatus 获取任务状态和结果
func GetTaskStatus(clt *core.SDKClient, auth *model.RequestHeader, taskRequest *report.GetTaskStatusRequest) (*model.ResponseHeader, []report.ReportTaskStatus, error) {
	req := &model.Request{
		Header: auth,
		Body:   taskRequest,
	}
	var resp report.GetTaskStatusResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
