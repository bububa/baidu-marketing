package model

// RequestHeader 请求header对象
type RequestHeader struct {
	// Username 推广账户名称
	Username string `json:"username,omitempty"`
	// Password 推广账户密码
	Password string `json:"password,omitempty"`
	// Token 您的token值
	Token string `json:"token"`
	// Target 被MCC账户管辖的普通推广账户名称
	Target string `json:"target,omitempty"`
	// AccessToken 百度商业服务市场服务商的access_token。注意属性名是大写的T
	AccessToken string `json:"access_token,omitempty"`
}

// RequestBody 请求业务数据
type RequestBody interface {
	Url() string // 接口链接
}

// Request API 请求对象
type Request struct {
	// Header header 对象
	Header RequestHeader `json:"header"`
	// Body 业务对象
	Body RequestBody `json:"body"`
}

// Url 请求API 地址
func (r Request) Url() string {
	return r.Body.Url()
}
