package image

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/asset/image"
)

// GetImage 通过图片规格、大小、格式、时间区间等条件，筛选出合适的图片返回
func GetImage(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *image.GetImageRequest) (*model.ResponseHeader, []image.Image, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image.GetImageResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
