package oauth

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// URLRequest 程序化拼接授权链接 API Request
type URLRequest struct {
	// AppID 需要授权的应用 ID
	AppID string `json:"appId,omitempty"`
	// Scope 应用权限代码，建议从应用管理界面系统生成的授权链接中获取。
	Scope string `json:"scope,omitempty"`
	// State 开发者自定义参数，长度限制 512 个字符，特殊字符需要 URLEncode
	State string `json:"state,omitempty"`
	// Callback 应用回调链接
	Callback string `json:"callback,omitempty"`
}

// URL 程序化拼接授权链接
func (r URLRequest) URL() string {
	values := util.GetUrlValues()
	defer util.PutUrlValues(values)
	values.Set("appId", r.AppID)
	values.Set("scope", r.Scope)
	values.Set("state", r.State)
	values.Set("callback", r.Callback)
	return util.StringsJoin(model.BASE_OAUTH_URL, "page/index?", values.Encode())
}
