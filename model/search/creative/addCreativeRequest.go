package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type AddCreativeRequest struct {
	CreativeTypes []Creative `json:"creativeTypes"` // 新增推广创意物料;集合长度限制：[1, 3000];建议分批多次请求
}

func (r AddCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeService/addCreative", model.BASE_URL_SMS)
}
