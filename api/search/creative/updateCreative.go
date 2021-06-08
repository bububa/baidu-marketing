package creative

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// 修改推广创意
func UpdateCreative(clt *core.SDKClient, auth model.RequestHeader, creatives []creative.Creative) ([]creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.UpdateCreativeRequest{
			CreativeTypes: creatives,
		},
	}
	var resp creative.UpdateCreativeResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
