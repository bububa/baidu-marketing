package newimage

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// UploadImageRequest 上传图片素材API Request
type UploadImageRequest struct {
	// Items 图片上传对象, 一次可以上传最多20张图片
	Items []UploadImageItem `json:"items"`
	// AddImage 是否保存到我的图库，默认为否。ture：设置为true，上传且保存，可在网页端推广平台-资产中心-我的图片查询；false：设置为否，仅上传不保存
	AddImage *bool `json:"addImage,omitempty"`
}

func (r UploadImageRequest) Url() string {
	return fmt.Sprintf("%sImageManageService/uploadImage", model.BASE_URL_SMS)
}

type UploadImageItem struct {
	// Content 图片的base64编码，不能为空，请求体最大不超过10M 。支持jpg/jpeg/png格式，注：请求时注意不要带"data:image/jpg;base64"。
	Content string `json:"content,omitempty"`
	// ImgMd5 图片的md5
	ImgMd5 string `json:"imgmd5,omitempty"`
	// ImageName 图片名称，不超过200个字符。英文字母占1个字符，1个汉字占2个字符
	ImageName string `json:"imageName,omitempty"`
}
