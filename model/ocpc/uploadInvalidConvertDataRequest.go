package ocpc

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// UploadInvalidConvertDataRequest
type UploadInvalidConvertDataRequest struct {
	// Token API接口token，每个推广账号对应一个唯一TOKEN
	Token string `json:"token"`
	// InvalidConversionTypes 回传转化数据数组，每次回传转化数据数组长度需小于100
	InvalidConversionTypes []InvalidConversionType `json:"invalidConversionTypes,omitempty"`
}

// InvalidConversionType
type InvalidConversionType struct {
	// LogidUrl logidUrl为带有&bd_vid=xxx的落地页url地址，只有在百度搜索广告里点击进入的落地页url，才会带有&bd_vid，如有相关疑问，可参考文档进行排查
	LogidUrl string `json:"logidUrl"`
	// ConvertType 转化类型，每次选择一个类型回传，模型可用的转化类型请以推广后台新建转化追踪披露的转化类型为准
	ConvertType int `json:"convertType"`
	// ConvertTime unix时间戳（精确到秒）;转化类型为46时必填
	ConvertTime int64 `json:"convertTime"`
	// Ip 无效转化的IP地址
	Ip string `json:"ip,omitempty"`
	// Confidence 置信度，0-100数字
	Confidence int `json:"confidence,omitempty"`
	// InvalidReason 无效转化描述
	InvalidReason string `json:"invalidReason,omitempty"`
}

func (r UploadInvalidConvertDataRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_OCPC, "uploadInvalidConvertData")
}

func (r UploadInvalidConvertDataRequest) OcpcToken() string {
	return r.Token
}

func (r *UploadInvalidConvertDataRequest) SetOcpcToken(token string) {
	r.Token = token
}
