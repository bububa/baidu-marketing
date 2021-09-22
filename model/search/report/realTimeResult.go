package report

import "encoding/json"

type RealTimeResult struct {
	ID          string        `json:"ID,omitempty"`          // 请求对象的ID app下载报告/推广电话报告，请求对象的ID，为null
	Name        []string      `json:"name,omitempty"`        // 请求对象的name，数组形式，不同报告形式的不同定义规则定义如下文
	RelatedId   int           `json:"relatedId,omitempty"`   // 依赖ID仅在 reportType=9，relatedId为adgroupId reportType=3，relatedId为userId reportType=5，relatedId为省市一级地域代码 其余情况为Null。
	Date        string        `json:"date,omitempty"`        // 统计开始时间
	KPIs        []json.Number `json:"KPIs,omitempty"`        // 按照请求顺序，返回KPI数据数组
	TotalNumber int64         `json:"totalNumber,omitempty"` // 记录总条数
	PageIndex   int           `json:"pageIndex,omitempty"`   // 当前页码
}
