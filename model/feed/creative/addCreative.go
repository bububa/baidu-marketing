package creative

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// AddCreativeRequest 新增推广创意 API Request
type AddCreativeRequest struct {
	// CreativeFeedTypes 新增信息流创意，单次最多添加100个创意
	CreativeFeedTypes []Creative `json:"creativeFeedTypes,omitempty"`
}

func (r AddCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CreativeFeedService/addCreativeFeed")
}

// AddCreativeResponse 新增推广创意 API Response
type AddCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
