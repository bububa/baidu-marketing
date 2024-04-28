package model

import (
	"encoding/json"
	"strconv"

	"github.com/xiaoshouchen/baidu-marketing/util"
)

// ResponseHeader API 返回header对象
type ResponseHeader struct {
	// Status 0:成功，1:部分失败，2:全部失败，3:系统错误
	Status int `json:"status,omitempty"`
	// Desc 描述
	Desc string `json:"desc,omitempty"`
	// Rquota 剩余的请求配额(现可忽略)
	Rquota int64 `json:"rquota,omitempty"`
	// Quota 本次请求发送的数据条数
	Quota int `json:"quota,omitempty"`
	// Failures  错误信息
	Failures []ResponseFailure `json:"failures,omitempty"`
	// Errors 错误信息
	Errors []ResponseFailure `json:"errors,omitempty"`
	// Oprs 成功操作数据条数
	Oprs int `json:"oprs,omitempty"`
	// Oprtime 操作时间描述
	Oprtime int64 `json:"oprtime,omitempty"`
}

// ResponseFailure 返回错误信息
type ResponseFailure struct {
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
	Position string `json:"position"`
}

// Response API 返回结果
type Response struct {
	// Header header对象
	Header ResponseHeader `json:"header,omitempty"`
	// Body 业务对象
	Body json.RawMessage `json:"body,omitempty"`
}

// IsError 判断结果是否错误
func (r Response) IsError() bool {
	return r.Header.Status != 0
}

// Error implement error interface
func (r Response) Error() string {
	buf, _ := json.Marshal(r.Header.Failures)
	return string(buf)
}

type DataResponse struct {
	// Code 请求处理状态，0：处理成功，非0：处理失败
	Code int `json:"code,omitempty"`
	// Message 请求处理话术
	Message string `json:"message,omitempty"`
	// Data 请求处理结果
	Data json.RawMessage `json:"data,omitempty"`
}

// IsError 判断结果是否错误
func (r DataResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r DataResponse) Error() string {
	return util.StringsJoin("code:", strconv.Itoa(r.Code), ", message:", r.Message)
}

// ActionCbResponse 转化追踪回调返回
type ActionCbResponse struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

func (r ActionCbResponse) IsError() bool {
	return r.ErrorCode != 0
}

func (r ActionCbResponse) Error() string {
	return r.ErrorMsg
}
