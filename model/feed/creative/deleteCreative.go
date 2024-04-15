package creative

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// DeleteCreativeRequest 删除推广创意 API Request
type DeleteCreativeRequest struct {
	// CreativeFeedIds 创意ID
	CreativeFeedIds []uint64 `json:"creativeFeedIds"`
}

func (r DeleteCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CreativeFeedService/deleteCreativeFeed")
}

// DeleteCreativeResponse 删除推广创意 API Response
type DeleteCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
