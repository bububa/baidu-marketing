package adgroup

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// AddAdgroupRequest 添加单元 API Request
type AddAdgroupRequest struct {
	// AdgroupFeedTypes 新增推广单元物料
	AdgroupFeedTypes []Adgroup `json:"adgroupFeedTypes"`
}

func (r AddAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AdgroupFeedService/addAdgroupFeed")
}

// AddAdgroupResponse 添加单元 API Response
type AddAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
