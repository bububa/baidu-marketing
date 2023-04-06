package report

type AdStats struct {
	Impression        int64   `json:"impression,omitempty"`        // Impression 表示展现次数。
	Click             int64   `json:"click,omitempty"`             // Click 表示点击次数。
	Cost              float64 `json:"cost,omitempty"`              // Cost 表示消费金额。
	CTR               float64 `json:"ctr,omitempty"`               // CTR 表示点击率。 计算公式：Click/Impression。
	CPC               float64 `json:"cpc,omitempty"`               // CPC 表示平均点击价格。 计算公式：Cost/Click。
	CPM               float64 `json:"cpm,omitempty"`               // CPM 表示千次展现消费金额。 计算公式：Cost/(Impression/1000)。
	PhoneButtonClicks int64   `json:"phoneButtonClicks,omitempty"` // PhoneButtonClicks 表示组件点击次数。 用于点击创意组件 button 次数。
	Interaction       int64   `json:"interaction,omitempty"`       // Interaction 表示广告的互动次数。 对于视频广告，互动行为包含点击互动和有效播放互动，对其他格式的广告，互动依然仅包含点击。
}
