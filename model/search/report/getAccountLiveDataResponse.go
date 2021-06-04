package report

type GetAccountLiveDataResponse struct {
	Data []AccountLiveData `json:"data,omitempty"`
}

type AccountLiveData struct {
	UpdateTime      string  `json:"updateTime,omitempty"`      // 数据更新时间; 默认返回： 日期类型，格式示例 "2017-08-23 15:00:00"
	UserId          int64   `json:"userId,omitempty"`          // 用户ID
	CampaignId      int64   `json:"campaignId,omitempty"`      // 计划ID
	TodayTotalClick int64   `json:"todayTotalClick,omitempty"` // 今日累计点击;根据请求数据类型不同，分别返回账户累计点击或计划累计点击
	TodayTotalCost  float64 `json:"todayTotalCost,omitempty"`  // 今日累计消费;根据请求数据类型不同，分别返回账户累计消费或计划累计消费
	TodayTotalCpc   float64 `json:"todayTotalCpc,omitempty"`   // 今日累计平均点击价格;根据请求数据类型不同，分别返回账户累计消费或计划累计平均点击价格
}
