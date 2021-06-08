package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type UpdateCreativeRequest struct {
	CreativeTypes []Creative `json:"creativeTypes,omitempty"` // 更新推广创意字段;集合长度限制：[1, 3000];建议分批多次请求
}

func (r UpdateCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeService/updateCreative", model.BASE_URL_SMS)
}
