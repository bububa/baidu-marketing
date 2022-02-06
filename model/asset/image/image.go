package image

// Image 图片素材
// Source 可选值
// 1：web
// 2：api
// 3：版权库
// 4：霓裳

// IsCollected 可选值
// 0：代表不收藏
// 1：代表收藏

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
	ImageId string `json:"imageId,omitempty"`
	// Url 图片上传成功后生成的链接
	Url string `json:"url,omitempty"`
	// Width 宽, 要求(宽_高)
	Width int `json:"width,omitempty"`
	// Height 高, 规格同上
	Height int `json:"height,omitempty"`
	// Size 图片大小, 1M以内
	Size int `json:"size,omitempty"`
	// Format 图片格式, jpg/jpeg/png
	Format string `json:"format,omitempty"`
	// Date 图片上传时间
	Date string `json:"date,omitempty"`
	// Desc 图片描述, 不超过200个字符;注：英文字母占1个字符，1个汉字占2个字符
	Desc string `json:"desc,omitempty"`
	// Signature 图片MD5值
	Signature string `json:"signature,omitempty"`
	// Source 图片来源
	Source int `json:"source,omitempty"`
	// IsCollected 图片收藏
	IsCollected *int `json:"isCollected,omitempty"`
}
