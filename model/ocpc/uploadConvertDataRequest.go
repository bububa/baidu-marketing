package ocpc

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

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
	ConvertValue int64 `json:"convertValue,omitempty"`
	// Confidence 置信度，0-100数字
	Confidence int `json:"confidence,omitempty"`
	// ExtInfo 点击&曝光原始信息，媒体侧进行数据拼接时使用
	ExtInfo string `json:"extInfo,omitempty"`
	// TheaterId 剧场id，广告主短剧剧场id
	TheaterId string `json:"theaterId,omitempty"`
	// TheaterShortPlayId 短剧id，广告主剧场下短剧的唯一id
	TheaterShortPlayId string `json:"theaterShortPlayId,omitempty"`
	// TheaterUserId 剧场内用户注册id，广告主剧场下的用户唯一标识
	TheaterUserId string `json:"theaterUserId,omitempty"`
	// OuterEventId 用户自定义的行为 id 标识，若上报有重复且需要去重，请填写该id，系统会根据该ID进行去重。最大值 255 字节
	OuterEventId string `json:"outerEventId,omitempty"`
	// OuterEventType 外部事件类型：当前需求有两类取值。quality_points_add 表示质量分累加。 custom_dedup 表示自定义去重。
	OuterEventType string `json:"outerEventType,omitempty"`
	// AttributeSource （转化来源：百度（0），自然流量（1），竞媒（2））
	AttributeSource int `json:"attributeSource,omitempty"`
	// InteractionsType （归因请求方式，点击(1), 播放(2), 关注(3), 分享(4), 点赞(5), 曝光（6）不传使用1，默认为点击归因
	InteractionsType int `json:"interactionsType,omitempty"`
}

func (r UploadConvertDataRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_OCPC, "uploadConvertData")
}

func (r UploadConvertDataRequest) OcpcToken() string {
	return r.Token
}

func (r *UploadConvertDataRequest) SetOcpcToken(token string) {
	r.Token = token
}
