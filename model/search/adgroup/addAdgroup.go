package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AddAdgroupRequest 添加单元 API Request
type AddAdgroupRequest struct {
	// AdgroupTypes 新增推广单元物料;集合长度限制：[1, 5000]
	AdgroupTypes []Adgroup `json:"adgroupTypes"`
}

func (r AddAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupService/addAdgroup")
}

// AddAdgroupResponse 添加单元 API Response
type AddAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
