package creative

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// 新增推广创意
func AddCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.AddCreativeRequest) ([]creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.AddCreativeResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
