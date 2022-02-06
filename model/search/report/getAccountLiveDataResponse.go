package report

// GetAccountLiveDataResponse 账户实时数据 API Response
type GetAccountLiveDataResponse struct {
	Data []AccountLiveData `json:"data,omitempty"`
}

// AccountLiveData 账户实时数据
type AccountLiveData struct {
	// UpdateTime 数据更新时间; 默认返回： 日期类型，格式示例 "2017-08-23 15:00:00"
	UpdateTime string `json:"updateTime,omitempty"`
	// UserId 用户ID
	UserId int64 `json:"userId,omitempty"`
	// CampaignId 计划ID
	CampaignId int64 `json:"campaignId,omitempty"`
	// TodayTotalClick 今日累计点击;根据请求数据类型不同，分别返回账户累计点击或计划累计点击
	TodayTotalClick int64 `json:"todayTotalClick,omitempty"`
	// TodayTotalCost 今日累计消费;根据请求数据类型不同，分别返回账户累计消费或计划累计消费
	TodayTotalCost float64 `json:"todayTotalCost,omitempty"`
	// TodayTotalCpc 今日累计平均点击价格;根据请求数据类型不同，分别返回账户累计消费或计划累计平均点击价格
	TodayTotalCpc float64 `json:"todayTotalCpc,omitempty"`
}
