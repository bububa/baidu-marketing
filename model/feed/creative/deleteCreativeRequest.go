package creative

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// DeleteCreativeRequest 删除推广创意 API Request
type DeleteCreativeRequest struct {
	// CreativeFeedIds 创意ID
	CreativeFeedIds []int64 `json:"creativeFeedIds"`
}

func (r DeleteCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeFeedService/deleteCreativeFeed", model.BASE_URL_FEED)
}
