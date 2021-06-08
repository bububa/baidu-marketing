package image

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
	ImageId     string `json:"imageId,omitempty"`     // 图片标识id
	Url         string `json:"url,omitempty"`         // 图片上传成功后生成的链接
	Width       int    `json:"width,omitempty"`       // 宽, 要求(宽_高)
	Height      int    `json:"height,omitempty"`      // 高, 规格同上
	Size        int    `json:"size,omitempty"`        // 图片大小, 1M以内
	Format      string `json:"format,omitempty"`      // 图片格式, jpg/jpeg/png
	Date        string `json:"date,omitempty"`        // 图片上传时间
	Desc        string `json:"desc,omitempty"`        // 图片描述, 不超过200个字符;注：英文字母占1个字符，1个汉字占2个字符
	Signature   string `json:"signature,omitempty"`   // 图片MD5值
	Source      int    `json:"source,omitempty"`      // 图片来源
	IsCollected *int   `json:"isCollected,omitempty"` // 图片收藏
}
