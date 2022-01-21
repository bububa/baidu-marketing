package report

import "encoding/json"

type RealTimeResult struct {
	ID             json.Number   `json:"ID,omitempty"`             // 请求对象的ID(ocpx报告、媒体ID报告，地域报告此字段无用，固定为0)
	Name           []string      `json:"name,omitempty"`           // 请求对象的name，数组形式，不同报告形式的不同定义规则见下文“name字段规则”
	Date           string        `json:"date,omitempty"`           // 统计开始时间
	KPIs           []json.Number `json:"KPIs,omitempty"`           // 按照请求顺序，返回KPI数据数组
	RelatedIdsList []int64       `json:"relatedIdsList,omitempty"` // 账户id，计划id，单元id
	TotalRowNumber json.Number   `json:"totalRowNumber,omitempty"` // 记录总条数
	PageIndex      int           `json:"pageindex,omitempty"`      // 当前页码
}
