package adgroup

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// DeleteAdgroupRequest 删除单元 API Request
type DeleteAdgroupRequest struct {
	// AdgroupIds 推广单元ID
	AdgroupIds []uint64 `json:"adgroupIds,omitempty"`
}

func (r DeleteAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupService/deleteAdgroup")
}
