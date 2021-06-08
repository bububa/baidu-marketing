package ocpc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

type ActionCbRequest struct {
	Akey        string `json:"akey,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	AType       string `json:"a_type,omitempty"`
	AValue      string `json:"a_value,omitempty"`
}

func (r ActionCbRequest) unsignedUrl() string {
	link := strings.Replace(r.CallbackUrl, "{{ATYPE}}", r.AType, -1)
	link = strings.Replace(link, "{{AVALUE}}", r.AValue, -1)
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
