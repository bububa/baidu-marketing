package model

// RequestHeader 请求header对象
type RequestHeader struct {
	// Username 推广账户名称
	Username string `json:"userName,omitempty"`
	// AccessToken 百度商业服务市场服务商的access_token。注意属性名是大写的T
	AccessToken string `json:"accessToken,omitempty"`
}

// RequestBody 请求业务数据
type RequestBody interface {
	Url() string // 接口链接
}

// Request API 请求对象
type Request struct {
	// Header header 对象
	Header *RequestHeader `json:"header"`
	// Body 业务对象
	Body RequestBody `json:"body"`
}

// SetUser set username password
func (r *Request) SetUser(username string, accessToken string) {
	if r.Header == nil {
		r.Header = new(RequestHeader)
	}
	r.Header.Username = username
	r.Header.AccessToken = accessToken
}

// Url 请求API 地址
func (r Request) Url() string {
	return r.Body.Url()
}

// ConversionRequest 转化请求
type ConversionRequest interface {
	RequestBody
	OcpcToken() string
	SetOcpcToken(token string)
}

type ActionCbRequest interface {
	RequestBody
}
