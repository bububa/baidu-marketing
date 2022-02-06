package ocpc

import (
	"encoding/json"

	"github.com/bububa/baidu-marketing/model"
)

// Response 转化回传返回结果
type Response struct {
	Header ResponseHeader `json:"header"`
}

func (r Response) IsError() bool {
	return r.Header.Status != 0
}

func (r Response) Error() string {
	buf, _ := json.Marshal(r.Header.Errors)
	return string(buf)
}

// ResponseHeader .
type ResponseHeader struct {
	Desc   string                  `json:"desc,omitempty"`
	Status int                     `json:"status,omitempty"` // 状态码
	Errors []model.ResponseFailure `json:"errors,omitempty"`
}
