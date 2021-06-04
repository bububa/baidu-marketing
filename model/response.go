package model

import "encoding/json"

type ResponseHeader struct {
	Status   int               `json:"status,omitempty"`   // //0:成功，1:部分失败，2:全部失败，3:系统错误
	Desc     string            `json:"desc,omitempty"`     // //描述
	Rquota   int64             `json:"rquota,omitempty"`   // //剩余的请求配额(现可忽略)
	Quota    int               `json:"quota,omitempty"`    // //本次请求发送的数据条数
	Failures []ResponseFailure `json:"failures,omitempty"` // //错误信息
	Oprs     int               `json:"oprs,omitempty"`     // //成功操作数据条数
	Oprtime  int64             `json:"oprtime,omitempty"`  // //操作时间描述
}

type ResponseFailure struct {
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
	Position string `json:"position"`
}

type Response struct {
	Header ResponseHeader  `json:"header,omitempty"`
	Body   json.RawMessage `json:"body,omitempty"`
}

func (r Response) IsError() bool {
	return r.Header.Status != 0
}

func (r Response) Error() string {
	buf, _ := json.Marshal(r.Header.Failures)
	return string(buf)
}
