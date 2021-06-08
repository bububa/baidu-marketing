package video

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
	Userid        int64  `json:"userid,omitempty"`        // 账户id
	Videoid       int64  `json:"videoid,omitempty"`       // 视频id
	Url           string `json:"url,omitempty"`           // 视频url
	VideoName     string `json:"videoName,omitempty"`     // 视频名称
	Source        int    `json:"source,omitempty"`        // 来源
	Capacity      int    `json:"capacity,omitempty"`      // 视频大小
	Format        string `json:"format,omitempty"`        // 视频格式
	Width         int    `json:"width,omitempty"`         // 宽度
	Height        int    `json:"height,omitempty"`        // 高度
	Duration      int    `json:"duration,omitempty"`      // 时长，单位s
	AddTime       int64  `json:"addTime,omitempty"`       // 添加时间，unix时间戳，单位毫秒
	ModTime       int64  `json:"modTime,omitempty"`       // 修改时间，unix时间戳，单位毫秒
	Thumbnail     string `json:"thumbnail,omitempty"`     // 封面图片url，为空代表该视频没有没有封面图，客户上传的视频肯定有
	DeliverStatus int    `json:"deliverStatus,omitempty"` // 投放状态，0：投放中，表示有创意正在引用该视频 1：未投放
	Istranscode   int    `json:"istranscode,omitempty"`   // 转码状态，0：转码成功，可绑定创意，1：转码中，可绑定创意，但需等转码成功之后才会审核，2：转码失败，不可绑定创意
	FromUserId    int64  `json:"fromUserId,omitempty"`    // 由谁分享，非分享的视频不返回此字段
	VideoMd5      string `json:"videoMd5,omitempty"`      // 视频文件md5
}
