package creative

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/creative"
)

func GetCreativeFeed(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.GetCreativeRequest) ([]creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.GetCreativeResponse
	err := clt.DoAny(req, &resp)
	if err != nil {
		return resp.Data, err //这里如果传入Ids为混杂的,会一部分正确一部分错误
	}
	return resp.Data, nil
}
