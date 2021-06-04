package ocpc

type UploadInvalidConvertDataRequest struct {
	Token                  string                  `json:"token"`                            // API接口token，每个推广账号对应一个唯一TOKEN
	InvalidConversionTypes []InvalidConversionType `json:"invalidConversionTypes,omitempty"` // 回传转化数据数组，每次回传转化数据数组长度需小于100
}

type InvalidConversionType struct {
	LogidUrl      string `json:"logidUrl"`                // logidUrl为带有&bd_vid=xxx的落地页url地址，只有在百度搜索广告里点击进入的落地页url，才会带有&bd_vid，如有相关疑问，可参考文档进行排查
	ConvertType   int    `json:"convertType"`             // 转化类型，每次选择一个类型回传，模型可用的转化类型请以推广后台新建转化追踪披露的转化类型为准
	ConvertTime   int64  `json:"convertTime"`             // unix时间戳（精确到秒）;转化类型为46时必填
	Ip            string `json:"ip,omitempty"`            // 无效转化的IP地址
	Confidence    int    `json:"confidence,omitempty"`    // 置信度，0-100数字
	InvalidReason string `json:"invalidReason,omitempty"` // 无效转化描述
}
