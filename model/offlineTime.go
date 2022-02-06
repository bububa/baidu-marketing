package model

// OfflineTime 推广下线时间
type OfflineTime struct {
	// Time 下线/上线时间点
	Time string `json:"time,omitempty"`
	// Flat 下线/上线状态: 1 - 上线; 0 - 下线
	Flat *int `json:"flat,omitempty"`
}
