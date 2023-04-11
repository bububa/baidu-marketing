package creative

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// UpdateCreativeRequest 修改推广创意 API Request
type UpdateCreativeRequest struct {
	// CreativeFeedTypes 更新信息流创意，单次最多更新100个创意
	CreativeFeedTypes []Creative `json:"creativeFeedTypes,omitempty"`
}

func (r UpdateCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeFeedService/updateCreativeFeed", model.BASE_URL_FEED)
}
