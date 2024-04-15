package creative

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateCreativeRequest 更新基础创意 API Request
type UpdateCreativeRequest struct {
	// CreativeTypes 更新推广创意字段;集合长度限制：[1, 3000];建议分批多次请求
	CreativeTypes []Creative `json:"creativeTypes,omitempty"`
}

func (r UpdateCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CreativeService/updateCreative")
}

// UpdateCreativeResponse 更新基础创意
type UpdateCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
