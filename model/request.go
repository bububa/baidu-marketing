package model

// 请求header对象
type RequestHeader struct {
	Username    string `json:"username,omitempty"`     // 推广账户名称
	Password    string `json:"password,omitempty"`     // 推广账户密码
	Token       string `json:"token"`                  // 您的token值
	Target      string `json:"target,omitempty"`       // 被MCC账户管辖的普通推广账户名称
	AccessToken string `json:"access_token,omitempty"` // 百度商业服务市场服务商的access_token。注意属性名是大写的T
}

type RequestBody interface {
	Url() string // 接口链接
}

type Request struct {
	Header RequestHeader `json:"header"`
	Body   RequestBody   `json:"body"`
}

func (r Request) Url() string {
	return r.Body.Url()
}
