package report

// Stat 统计
type Stat struct {
	Dimension
	CommonMetric
	FeedMetric
	SearchMetric
}

// CommonMetric 通用指标
type CommonMetric struct {
	// Impression 表示展现次数。
	Impression int64 `json:"impression,omitempty"`
	// Click 表示点击次数。
	Click int64 `json:"click,omitempty"`
	// Cost 表示消费金额。
	Cost float64 `json:"cost,omitempty"`
	// CTR 表示点击率。 计算公式：Click/Impression。
	CTR float64 `json:"ctr,omitempty"`
	// CPC 表示平均点击价格。 计算公式：Cost/Click。
	CPC float64 `json:"cpc,omitempty"`
	// CPM 表示千次展现消费金额。 计算公式：Cost/(Impression/1000)。
	CPM float64 `json:"cpm,omitempty"`
}
