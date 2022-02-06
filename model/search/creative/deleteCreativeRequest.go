package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// DeleteCreativeRequest 删除推广创意 API Request
type DeleteCreativeRequest struct {
	// CreativeIds 创意ID
	CreativeIds []int64 `json:"creativeIds"`
}

func (r DeleteCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeService/deleteCreative", model.BASE_URL_SMS)
}
