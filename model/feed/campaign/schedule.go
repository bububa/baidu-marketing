package campaign

// Schedule 暂停时段设置
type Schedule struct {
	// WeekDay 推广暂停日。取值范围如下：1：星期一2：星期二3：星期三4：星期四5：星期五6：星期六7：星期日
	WeekDay int `json:"weekDay,omitempty"`
	// StartHour 推广暂停时段开始时间，以小时为单位，取值范围:[0,23]
	StartHour int `json:"startHour,omitempty"`
	// EndHour 推广暂停时段结束时间，以小时为单位，取值范围:[1,24]
	EndHour int `json:"endHour,omitempty"`
}
