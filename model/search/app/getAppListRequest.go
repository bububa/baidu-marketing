package app

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

type GetAppListRequest struct {
	Platforms []int `json:"platforms,omitempty"`
	Limit     []int `json:"limit,omitempty"`
}

func (r GetAppListRequest) Url() string {
	return fmt.Sprintf("%sAppProcessService/getAppList", model.BASE_URL_SMS)
}
