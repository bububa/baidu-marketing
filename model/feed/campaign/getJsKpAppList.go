package campaign

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetJsKpAppListRequest 查询APP信息 API Request
type GetJsKpAppListRequest struct {
}

func (r GetJsKpAppListRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AppFeedService/getJsKpAppList")
}

// GetJsKpAppListResponse 查询APP信息 API Response
type GetJsKpAppListResponse struct {
	Data []AppInfo `json:"data,omitempty"`
}
