package image

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/asset/image"
)

// UploadImage 图片上传接口，上传后搜索推广与信息流推广可共用。上传时不限制单张图片大小，单次请求不超过10M。
func UploadImage(clt *core.SDKClient, auth model.RequestHeader, reqBody *image.UploadImageRequest) (*model.ResponseHeader, []image.Image, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image.UploadImageResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
