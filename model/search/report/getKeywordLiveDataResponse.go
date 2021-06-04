package report

type GetKeywordLiveDataResponse struct {
	Data []KeywordLiveData `json:"data,omitempty"`
}

type KeywordLiveData struct {
	KeywordId   int64   `json:"keywordId,omitempty"`   // 关键词Id
	RegionId    int     `json:"regionId,omitempty"`    // 地域代码
	Device      int     `json:"device,omitempty"`      // 投放设备类型
	Minute      string  `json:"minute,omitempty"`      // 数据时间
	LeftRank    float64 `json:"leftRank,omitempty"`    // 关键词在minute指定的分钟内左侧平均排名（计算机）
	RightRank   float64 `json:"rightRank,omitempty"`   // 关键词在minute指定的分钟内右侧平均排名（计算机）
	TopRank     float64 `json:"topRank,omitempty"`     // 关键词在minute指定的分钟内上方平均排名（移动设备）
	BottomRank  float64 `json:"bottomRank,omitempty"`  // 关键词在minute指定的分钟内下方平均排名（移动设备）
	LeftShows   int64   `json:"leftShows,omitempty"`   // 关键词在minute指定的分钟内左侧展现次数（计算机）
	RightShows  int64   `json:"rightShows,omitempty"`  // 关键词在minute指定的分钟内右侧展现次数（计算机）
	TopShows    int64   `json:"topShows,omitempty"`    // 关键词在minute指定的分钟内上方展现次数（移动设备)
	BottomShows int64   `json:"bottomShows,omitempty"` // 关键词在minute指定的分钟内下方展现次数（移动设备）
	Click       int64   `json:"click,omitempty"`       // 关键词在minute指定的分钟内的点击次数
	Cost        float64 `json:"cost,omitempty"`        // 关键词在minute指定的分钟内的消费
}
