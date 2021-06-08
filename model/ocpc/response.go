package ocpc

import (
	"encoding/json"

	"github.com/bububa/baidu-marketing/model"
)

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

type ResponseHeader struct {
	Desc   string                  `json:"desc,omitempty"`
	Status int                     `json:"status,omitempty"` // 状态码
	Errors []model.ResponseFailure `json:"errors,omiempty"`
}
