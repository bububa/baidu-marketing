package report

import "encoding/json"

// RealTimeResult 实时报告
type RealTimeResult struct {
	// ID 请求对象的ID(ocpx报告、媒体ID报告，地域报告此字段无用，固定为0)
	ID json.Number `json:"ID,omitempty"`
	// Name 请求对象的name，数组形式，不同报告形式的不同定义规则见下文“name字段规则”
	Name []string `json:"name,omitempty"`
	// Date 统计开始时间
	Date string `json:"date,omitempty"`
	// KPIs 按照请求顺序，返回KPI数据数组
	KPIs []json.Number `json:"KPIs,omitempty"`
	// RelatedIDsList 账户id，计划id，单元id
	RelatedIdsList []int64 `json:"relatedIdsList,omitempty"`
	// RelateIDsList 账户id，计划id，单元id
	RelateIdsList []json.Number `json:"relateIdsList,omitempty"`
	// TotalRowNumber 记录总条数
	TotalRowNumber json.Number `json:"totalRowNumber,omitempty"`
	// PageIndex 当前页码
	PageIndex json.Number `json:"pageindex,omitempty"`
}
