package ocpc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// ActionCbRequest 转化追踪回调 请求
type ActionCbRequest struct {
	Akey        string `json:"akey,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	AType       string `json:"a_type,omitempty"`
	AValue      string `json:"a_value,omitempty"`
}

func (r ActionCbRequest) unsignedUrl() string {
	link := strings.ReplaceAll(r.CallbackUrl, "{{ATYPE}}", r.AType)
	link = strings.ReplaceAll(link, "{{AVALUE}}", r.AValue)
	return link
}

func (r ActionCbRequest) Sign() string {
	signUrl := fmt.Sprintf("%s%s", r.unsignedUrl(), r.Akey)
	h := md5.New()
	h.Write([]byte(signUrl))
	return hex.EncodeToString(h.Sum(nil))
}

func (r ActionCbRequest) Url() string {
	return fmt.Sprintf("%s&sign=%s", r.unsignedUrl(), r.Sign())
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
