package model

// TimeId 取值为3位整数，从左至右:
// 第一位表示每周的星期几，取值为 1-7
// 第二三位表示小时编号，取值范围为 00-23
// 不设置的小时内将不投放
// 例如：设置每周一的0点到1点投放时，该值取100；每周六的22点到23点投放时，该值取622
type SchedulePriceFactor struct {
	TimeId      int     `json:"timeId,omitempty"`      // 时间段编号
	PriceFactor float64 `json:"priceFactor,omitempty"` // 出价系数;取值范围：[0.1, 10.0]
}
