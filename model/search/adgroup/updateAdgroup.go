package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// UpdateAdgroupRequest 更新推广单元 API Request
type UpdateAdgroupRequest struct {
	// AdgroupTypes 更新推广单元字段;集合长度限制：[1, 5000]
	AdgroupTypes []Adgroup `json:"adgroupTypes,omitempty"`
}

func (r UpdateAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupService/updateAdgroup")
}

// UpdateAdgroupResponse 更新推广单元API Response
type UpdateAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
