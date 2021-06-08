package creative

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// 查询推广创意
func GetCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.GetCreativeRequest) ([]creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.GetCreativeResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
