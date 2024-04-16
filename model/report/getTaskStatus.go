package report

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetTaskStatusRequest 获取异步任务状态 API Request
type GetTaskStatusRequest struct {
	// TaskId 异步任务ID，任务ID的有效期是10分钟。
	TaskId string `json:"taskId,omitempty"`
}

func (r GetTaskStatusRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "OpenApiReportService/getTaskStatus")
}

// GetTaskStatusResponse 获取异步任务状态 API Response
type GetTaskStatusResponse struct {
	Data []ReportTaskStatus `json:"data,omitempty"`
}

type ReportTaskStatus struct {
	// TaskStatus 任务状态。
	// SUBMITTED：已提交
	// RUNNING：运行中
	// SUCCESS：成功
	// FAIL：失败
	TaskStatus string `json:"taskStatus,omitempty"`
	// SubmitTime 任务提交的时间
	SubmitTime string `json:"submitTime,omitempty"`
	// StartTime 任务开始执行的时间
	StartTime string `json:"startTime,omitempty"`
	// FinishTime 任务完成的时间
	FinishTime string `json:"finishTime,omitempty"`
	// FileUrl 文件地址，当taskStatus为成功时，fileUrl为文件地址，其他状态为空。
	// 地址的有效期为5分钟
	FileUrl string `json:"fileUrl,omitempty"`
	// DataStartRow 数据在文件中的起始行。
	// 文件的前几行为表头或其他信息，dataStartRow为数据的起始行。
	DataStartRow int64 `json:"dataStartRow,omitempty"`
	// TableHeader 数据每一列对应的字段名。
	TableHeader []string `json:"tableHeader,omitempty"`
}
