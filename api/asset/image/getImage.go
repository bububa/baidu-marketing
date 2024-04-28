package image

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/asset/image"
)

// GetImage 通过图片规格、大小、格式、时间区间等条件，筛选出合适的图片返回
func GetImage(clt *core.SDKClient, auth *model.RequestHeader, reqBody *image.GetImageRequest) (*model.ResponseHeader, []image.Image, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image.GetImageResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
