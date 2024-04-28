package app

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AddAdgroupAppBindRequest 添加APP绑定 API Request
type AddAdgroupAppBindRequest struct {
	// AppBinds 新增APP绑定对象数组 单次请求不超过2000个
	AppBinds []AppBindItem `json:"appBinds,omitempty"`
}

func (r AddAdgroupAppBindRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupAppService/addAdgroupAppBind")
}

// AddAdgroupAppBindResponse 添加APP绑定 API Response
type AddAdgroupAppBindResponse struct {
	Data []AppBindItem `json:"data,omitempty"`
}
