package adgroup

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateAdgroupRequest 更新推广单元 API Request
type UpdateAdgroupRequest struct {
	// AdgroupTypes 更新推广单元字段;集合长度限制：[1, 5000]
	AdgroupFeedTypes []Adgroup `json:"adgroupFeedTypes,omitempty"`
}

func (r UpdateAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AdgroupFeedService/updateAdgroupFeed")
}

// UpdateAdgroupResponse 更新推广单元API Response
type UpdateAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
