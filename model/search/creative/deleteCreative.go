package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// DeleteCreativeRequest 删除基础创意 API Request
type DeleteCreativeRequest struct {
	// CreativeIds 创意ID
	CreativeIds []uint64 `json:"creativeIds"`
}

func (r DeleteCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CreativeService/deleteCreative")
}
