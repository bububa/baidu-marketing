package report

import (
	"github.com/xiaoshouchen/baidu-marketing/enum"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// CreateReportTaskRequest 创建异步任务 API Request
type CreateReportTaskRequest struct {
	// ReportType 报告类型，报告唯一标识，点此查看全部推广报告reportType
	ReportType enum.ReportType `json:"reportType"`
	// UserIds 查询的用户ID，用于查询超管账户所管辖的多个账户数据。可为空，默认只查询当前账户数据。
	UserIds []string `json:"userIds,omitempty"`
	// TimeUnit 时间单位，每个报告支持的时间单位见报告说明文档。
	// HOUR: 小时
	// DAY: 天
	// WEEK: 周
	// MONTH: 月
	// SUMMARY: 时间段汇总
	TimeUnit enum.TimeUnit `json:"timeUnit,omitempty"`
	// StartDate 数据的起始日期，格式 2020-05-28
	StartDate string `json:"startDate,omitempty"`
	// EndDate 数据的结束日期，格式 2020-05-29
	EndDate string `json:"endDate,omitempty"`
	// Columns 	查询的列，包含属性和转化指标，必填项，至少要带一个转化指标。
	// 每个报告的说明文档里都有支持的columns列表及说明。
	Columns []string `json:"columns,omitempty"`
	// Sorts 排序信息，详见下方排序说明
	Sorts []Sort `json:"sorts,omitempty"`
	// Filters 筛选条件集合，非必填，详见下方过滤条件说明。
	Filters []Filter `json:"filters,omitempty"`
	// NeedSum 	是否需总计，非必填，默认不需要总计。
	NeedSum bool `json:"needSum,omitempty"`
}

func (r CreateReportTaskRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "OpenApiReportService/createReportTask")
}

// CreateReportTaskResponse 创建异步任务 API Response
type CreateReportTaskResponse struct {
	Data struct {
		// TaskId 异步任务ID，任务ID有效期为10分钟。
		TaskId string `json:"taskId,omitempty"`
	} `json:"data,omitempty"`
}
