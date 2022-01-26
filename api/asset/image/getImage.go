package image

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/asset/image"
)

// 通过图片规格、大小、格式、时间区间等条件，筛选出合适的图片返回
func GetImage(clt *core.SDKClient, auth model.RequestHeader, reqBody *image.GetImageRequest) ([]image.Image, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image.GetImageResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
