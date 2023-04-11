package creative

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// AddCreativeRequest 新增推广创意 API Request
type AddCreativeRequest struct {
	// CreativeFeedTypes 新增信息流创意，单次最多添加100个创意
	CreativeFeedTypes []Creative `json:"creativeFeedTypes"`
}

func (r AddCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeFeedService/addCreativeFeed", model.BASE_URL_FEED)
}
