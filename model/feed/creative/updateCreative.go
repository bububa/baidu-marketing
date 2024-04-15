package creative

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateCreativeRequest 修改推广创意 API Request
type UpdateCreativeRequest struct {
	// CreativeFeedTypes 更新信息流创意，单次最多更新100个创意
	CreativeFeedTypes []Creative `json:"creativeFeedTypes,omitempty"`
}

func (r UpdateCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CreativeFeedService/updateCreativeFeed")
}

// UpdateCreativeResponse 修改推广创意
type UpdateCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
