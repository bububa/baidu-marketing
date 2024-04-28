package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/creative"
)

// UpdateCreative 修改推广创意
func UpdateCreative(clt *core.SDKClient, auth *model.RequestHeader, creatives []creative.Creative) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.UpdateCreativeRequest{
			CreativeFeedTypes: creatives,
		},
	}
	var resp creative.UpdateCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
