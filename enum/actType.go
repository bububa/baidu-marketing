package enum

// ActType 请求类型
type ActType int

const (
	// ActType_CLICK 点击请求
	ActType_CLICK ActType = 2
	// ActType_IMPRESSION 曝光请求
	ActType_IMPRESSION ActType = 3
)
