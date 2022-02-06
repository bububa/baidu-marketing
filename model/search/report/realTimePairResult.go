package report

import "encoding/json"

// RealTimePairResult 配对报告返回
type RealTimePairResult struct {
	// KeywordId 关键词ID
	KeywordId int64 `json:"keywordId,omitempty"`
	// CreativeId 创意ID
	CreativeId int64 `json:"creativeId,omitempty"`
	// PairInfo 账户名，计划名，单元名，关键词字面，创意标题,创意描述一，创意描述二，创意显示url ，搜索引擎，精确匹配扩展(地域词扩展)触发
	PairInfo []string `json:"pairInfo,omitempty"`
	// Date 统计开始时间
	Date string `json:"date,omitempty"`
	// KPIs 按照请求顺序，返回KPI数据数组
	KPIs []json.Number `json:"KPIs,omitempty"`
}
