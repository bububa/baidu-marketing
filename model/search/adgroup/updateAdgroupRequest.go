package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// UpdateAdgroupRequest 更新推广单元 API Request
type UpdateAdgroupRequest struct {
	// AdgroupTypes 更新推广单元字段;集合长度限制：[1, 5000]
	AdgroupTypes []Adgroup `json:"adgroupTypes,omitempty"`
}

func (r UpdateAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/updateAdgroup", model.BASE_URL_SMS)
}
