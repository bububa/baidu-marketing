package video

// Video 视频素材
// Source 可选值
// 1：视频平台
// 2：信息流视频
// 7：百度推广客户端
// 8：推荐视频
// 10:创意中心-自提
// 11：聚屏
// 18：多镜头模板制作
// 19：慧合平台
// 20：CIP视频素材库
// 21：视频片段
// 23：慧视自定义模板合成
// 24：信息流站点挖掘
// 25：信息流尺寸转置
// 31：慧拍
// 32：视频编辑器

// DeliverStatus 可选值
// 0：投放中，表示有创意正在引用该视频
// 1：未投放

// Istranscode 可选值
// 0：转码成功，可绑定创意
// 1：转码中，可绑定创意，但需等转码成功之后才会审核
// 2：转码失败，不可绑定创意
type Video struct {
	// Userid  账户id
	Userid int64 `json:"userid,omitempty"`
	// Videoid 视频id
	Videoid int64 `json:"videoid,omitempty"`
	// Url 视频url
	Url string `json:"url,omitempty"`
	// VideoName 视频名称
	VideoName string `json:"videoName,omitempty"`
	// Source 来源
	Source int `json:"source,omitempty"`
	// Capacity 视频大小
	Capacity int `json:"capacity,omitempty"`
	// Format 视频格式
	Format string `json:"format,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
	// Duration 时长，单位s
	Duration int `json:"duration,omitempty"`
	// AddTime 添加时间，unix时间戳，单位毫秒
	AddTime int64 `json:"addTime,omitempty"`
	// ModTime 修改时间，unix时间戳，单位毫秒
	ModTime int64 `json:"modTime,omitempty"`
	// Thumbnail 封面图片url，为空代表该视频没有没有封面图，客户上传的视频肯定有
	Thumbnail string `json:"thumbnail,omitempty"`
	// DeliverStatus 投放状态，0：投放中，表示有创意正在引用该视频 1：未投放
	DeliverStatus int `json:"deliverStatus,omitempty"`
	// Istranscode 转码状态，0：转码成功，可绑定创意，1：转码中，可绑定创意，但需等转码成功之后才会审核，2：转码失败，不可绑定创意
	Istranscode int `json:"istranscode,omitempty"`
	// FromUserId 由谁分享，非分享的视频不返回此字段
	FromUserId int64 `json:"fromUserId,omitempty"`
	// VideoMd5 视频文件md5
	VideoMd5 string `json:"videoMd5,omitempty"`
}
