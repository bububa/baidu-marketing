package trans

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateOcpcTransFeedRequest 更新转化追踪 API Request
type UpdateOcpcTransFeedRequest struct {
	// TransTraceUpdateTypes 转化追踪数据对象
	TransTraceUpdateTypes []OcpcTransFeed `json:"transTraceUpdateTypes,omitempty"`
}

func (r UpdateOcpcTransFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "TraceApiService/updateTransTrace")
}

// UpdateOcpcTransFeedResponse 修改转化追踪 API Response
type UpdateOcpcTransFeedResponse struct {
	Data []OcpcTransFeed `json:"data,omitempty"`
}
