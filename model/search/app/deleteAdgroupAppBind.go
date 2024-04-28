package app

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// DeleteAdgroupAppBindRequest 删除APP绑定 API Request
type DeleteAdgroupAppBindRequest struct {
	// DelBindIds 删除绑定Id信息
	// 单次请求不超过2000个
	DelBindIds []uint64 `json:"delBindIds,omitempty"`
}

func (r DeleteAdgroupAppBindRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupAppService/deleteAdgroupAppBind")
}

// DeleteAdgroupAppBindResponse 删除APP绑定 API Response
type DeleteAdgroupAppBindResponse struct {
	Data []AppBindItem `json:"data,omitempty"`
}
