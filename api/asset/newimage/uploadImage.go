package newimage

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	image2 "github.com/bububa/baidu-marketing/model/asset/newimage"
)

// UploadImage 图片上传接口，上传后搜索推广与信息流推广可共用。上传时不限制单张图片大小，单次请求不超过10M。
func UploadImage(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *image2.UploadImageRequest) (*model.ResponseHeader, []image2.Image, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image2.UploadImageResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
