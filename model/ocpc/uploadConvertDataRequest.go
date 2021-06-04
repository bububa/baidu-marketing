package ocpc

type UploadConvertDataRequest struct {
	Token           string           `json:"token"`                     // API接口token，每个推广账号对应一个唯一TOKEN
	ConversionTypes []ConversionType `json:"conversionTypes,omitempty"` // 回传转化数据数组，每次回传转化数据数组长度需小于100
}

type ConversionType struct {
	LogidUrl     string `json:"logidUrl"`               // logidUrl为带有&bd_vid=xxx的落地页url地址，只有在百度搜索广告里点击进入的落地页url，才会带有&bd_vid，如有相关疑问，可参考文档进行排查
	NewType      int    `json:"newType"`                // 转化类型，每次选择一个类型回传，模型可用的转化类型请以推广后台新建转化追踪披露的转化类型为准
	DeviceType   int    `json:"deviceType,omitempty"`   // 0代表安卓;1代表IOS;2代表其他
	DeviceId     string `json:"deviceId,omitempty"`     // 安卓:IMEI 号取；md5 sum 摘要；IOS:IDFA 号原值
	IsConvert    *int   `json:"isConvert,omitempty"`    // 是否发生转化（0或1）；0：未发生转化；1：发生转化; 默认为1
	ConvertTime  int64  `json:"convertTime,omitempty"`  // unix时间戳（精确到秒）;转化类型为46时必填
	ConvertValue int    `json:"convertValue,omitempty"` // 转化金额（单位分）
	Confidence   int    `json:"confidence,omitempty"`   // 置信度，0-100数字
}
