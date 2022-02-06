package report

// GetKeywordLiveDataResponse 关键词实时数据 API Response
type GetKeywordLiveDataResponse struct {
	Data []KeywordLiveData `json:"data,omitempty"`
}

// KeywordLiveData 关键词实时数据
type KeywordLiveData struct {
	// KeywordId 关键词Id
	KeywordId int64 `json:"keywordId,omitempty"`
	// RegionId 地域代码
	RegionId int `json:"regionId,omitempty"`
	// Device 投放设备类型
	Device int `json:"device,omitempty"`
	// Minute 数据时间
	Minute string `json:"minute,omitempty"`
	// LeftRank 关键词在minute指定的分钟内左侧平均排名（计算机）
	LeftRank float64 `json:"leftRank,omitempty"`
	// RightRank 关键词在minute指定的分钟内右侧平均排名（计算机）
	RightRank float64 `json:"rightRank,omitempty"`
	// TopRank 关键词在minute指定的分钟内上方平均排名（移动设备）
	TopRank float64 `json:"topRank,omitempty"`
	// BottomRank 关键词在minute指定的分钟内下方平均排名（移动设备）
	BottomRank float64 `json:"bottomRank,omitempty"`
	// LeftShows 关键词在minute指定的分钟内左侧展现次数（计算机）
	LeftShows int64 `json:"leftShows,omitempty"`
	// RightShows 关键词在minute指定的分钟内右侧展现次数（计算机）
	RightShows int64 `json:"rightShows,omitempty"`
	// TopShows 关键词在minute指定的分钟内上方展现次数（移动设备)
	TopShows int64 `json:"topShows,omitempty"`
	// BottomShows 关键词在minute指定的分钟内下方展现次数（移动设备）
	BottomShows int64 `json:"bottomShows,omitempty"`
	// Click 关键词在minute指定的分钟内的点击次数
	Click int64 `json:"click,omitempty"`
	// Cost 关键词在minute指定的分钟内的消费
	Cost float64 `json:"cost,omitempty"`
}
