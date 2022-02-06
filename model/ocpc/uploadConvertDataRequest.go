package ocpc

// UploadConvertDataRequest 广告主回传转化数据 API Request
type UploadConvertDataRequest struct {
	// Token API接口token，每个推广账号对应一个唯一TOKEN
	Token string `json:"token"`
	// ConversionTypes 回传转化数据数组，每次回传转化数据数组长度需小于100
	ConversionTypes []ConversionType `json:"conversionTypes,omitempty"`
}

// ConversionType 回转转化数据
type ConversionType struct {
	// LogidUrl logidUrl为带有&bd_vid=xxx的落地页url地址，只有在百度搜索广告里点击进入的落地页url，才会带有&bd_vid，如有相关疑问，可参考文档进行排查
	LogidUrl string `json:"logidUrl"`
	// NewType 转化类型，每次选择一个类型回传，模型可用的转化类型请以推广后台新建转化追踪披露的转化类型为准
	NewType int `json:"newType"`
	// DeviceType 0代表安卓;1代表IOS;2代表其他
	DeviceType int `json:"deviceType,omitempty"`
	// DeviceId 安卓:IMEI 号取；md5 sum 摘要；IOS:IDFA 号原值
	DeviceId string `json:"deviceId,omitempty"`
	// IsConvert 是否发生转化（0或1）；0：未发生转化；1：发生转化; 默认为1
	IsConvert *int `json:"isConvert,omitempty"`
	// ConvertTime unix时间戳（精确到秒）;转化类型为46时必填
	ConvertTime int64 `json:"convertTime,omitempty"`
	// ConvertValue 转化金额（单位分）
	ConvertValue int `json:"convertValue,omitempty"`
	// Confidence 置信度，0-100数字
	Confidence int `json:"confidence,omitempty"`
}
