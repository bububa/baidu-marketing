package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type UpdateAdgroupRequest struct {
	AdgroupTypes []Adgroup `json:"adgroupTypes,omitempty"` // 更新推广单元字段;集合长度限制：[1, 5000]
}

func (r UpdateAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/updateAdgroup", model.BASE_URL_SMS)
}
