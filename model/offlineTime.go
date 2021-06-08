package model

type OfflineTime struct {
	Time string `json:"time,omitempty"` // 下线/上线时间点
	Flat *int   `json:"flat,omitempty"` // 下线/上线状态: 1 - 上线; 0 - 下线
}
