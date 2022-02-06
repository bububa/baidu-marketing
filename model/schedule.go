package model

// Sechedule 投放排期
// WeekDay 可选值
// 1 - 星期一
// 2 - 星期二
// 3 - 星期三
// 4 - 星期四
// 5 - 星期五
// 6 - 星期六
// 7 - 星期日
type Schedule struct {
	// StartHour 开始时间; 以小时为单位，取值范围：[0,23]
	StartHour int `json:"startHour"`
	// EndHour 结束时间; 以小时为单位，取值范围：[1,24]
	EndHour int `json:"endHour"`
	// WeekDay 星期几
	WeekDay int `json:"weekDay"`
}
