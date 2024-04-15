package adgroup

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// DeleteAdgroupRequest 删除单元 API Request
type DeleteAdgroupRequest struct {
	// AdgroupFeedIds 推广单元ID
	AdgroupFeedIds []uint64 `json:"adgroupFeedIds"`
}

func (r DeleteAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AdgroupFeedService/deleteAdgroupFeed")
}

type DeleteAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
