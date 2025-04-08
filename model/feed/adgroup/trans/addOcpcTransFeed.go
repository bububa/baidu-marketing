package trans

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// AddOcpcTransFeedRequest 添加转化追踪 API Request
type AddOcpcTransFeedRequest struct {
	// OcpcTransFeedTypes 新增转化追踪
	OcpcTransFeedTypes []OcpcTransFeed `json:"ocpcTransFeedTypes,omitempty"`
}

func (r AddOcpcTransFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "OcpcTransFeedService/addOcpcTransFeed")
}

// AddOcpcTransFeedResponse 添加转化追踪 API Response
type AddOcpcTransFeedResponse struct {
	Data []OcpcTransFeed `json:"data,omitempty"`
}
