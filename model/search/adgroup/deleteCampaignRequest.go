package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// DeleteAdgroupRequest 删除单元 API Request
type DeleteAdgroupRequest struct {
	// AdgroupIds 推广单元ID
	AdgroupIds []int64 `json:"adgroupIds"`
}

func (r DeleteAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/deleteAdgroup", model.BASE_URL_SMS)
}
