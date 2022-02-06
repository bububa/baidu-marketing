package app

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetAppListRequest 获取APP素材 API Request
type GetAppListRequest struct {
	// Platforms app操作系统
	Platforms []int `json:"platforms,omitempty"`
	// Limit 分页信息，第一个参数偏移量，第二个参数是页面大小
	Limit []int `json:"limit,omitempty"`
}

func (r GetAppListRequest) Url() string {
	return fmt.Sprintf("%sAppProcessService/getAppList", model.BASE_URL_SMS)
}
