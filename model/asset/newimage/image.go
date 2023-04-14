package newimage

// Image 图片素材（新图片服务（ImageManageService）打通了搜索图库与信息流图库）
// SourceType 可选值
//0-本地上传
//1-本地上传
//3-图片工具
//4-图片工具
//5-图片工具
//6-图片工具
//7-图片工具
//8-图片工具
//14-图片工具
//16-图片工具
//9-版权图片
//10-站点挖掘
//11-视频抽帧

// 宽, 要求(宽_高)
// 370px_245px
// 1116px_627px
// 644px_280px
// 600px_248px
// 190px_190px
// 220px_220px
// 629px_90px
// 120px_90px
// 560px_170px
// 218px_146px
// 120px_120px

type Image struct {
	// ImageId 图片标识id
	ImageId int64 `json:"imageId,omitempty"`
	// Width 宽, 要求(宽_高)
	Width int `json:"width,omitempty"`
	// Height 高, 规格同上
	Height int `json:"height,omitempty"`
	// Size 图片大小, 1M以内
	Size int `json:"size,omitempty"`
	// Url 图片上传成功后生成的链接
	Url string `json:"url,omitempty"`
	// Format 图片格式, jpg/jpeg/png
	Format string `json:"format,omitempty"`
	// Signature 图片MD5值
	Signature string `json:"signature,omitempty"`
	// Date 添加时间
	AddTime string `json:"addTime,omitempty"`
	// Desc 修改时间
	ModTime string `json:"modTime,omitempty"`
	// IsCollected 图片收藏
	IsCollected *bool `json:"isCollected,omitempty"`
	// Source 图片来源类型
	SourceType int `json:"sourceType,omitempty"`
	// ImageName 图片名称;只有有图片名称的才会返回，没有图片名称（即名称为空）时不会返回 。如上传图片时设置该字段将会返回
	ImageName string `json:"imageName,omitempty"`
}
