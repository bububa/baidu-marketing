package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AddCreativeRequest 添加基础创意 API Request
type AddCreativeRequest struct {
	// CreativeTypes 新增推广创意物料;集合长度限制：[1, 3000];建议分批多次请求
	CreativeTypes []Creative `json:"creativeTypes"`
}

func (r AddCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CreativeService/addCreative")
}

// AddCreativeResponse 添加基础创意 API Response
type AddCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
