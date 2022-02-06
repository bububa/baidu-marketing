package report

import "encoding/json"

// RealTimeQueryResult 搜索词报告返回
type RealTimeQueryResult struct {
	// Query 搜索词字面
	Query string `json:"query,omitempty"`
	// KeywordId 关键词ID
	KeywordId int64 `json:"keywordId,omitempty"`
	// CreativeId 创意ID
	CreativeId int64 `json:"creativeId,omitempty"`
	// QueryInfo 账户名，计划名，单元名，关键词字面，创意标题,创意描述一，创意描述二，创意显示url ，搜索引擎，精确匹配扩展(地域词扩展)触发
	QueryInfo []string `json:"queryInfo,omitempty"`
	// Date 统计开始时间
	Date string `json:"Dat,omitempty"`
	// KPIs 按照请求顺序，返回KPI数据数组
	KPIs []json.Number `json:"KPIs,omitempty"`
}
