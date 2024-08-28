package creative

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// UpdateCreative 更新基础创意
func UpdateCreative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, creatives []creative.Creative) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: &creative.UpdateCreativeRequest{
			CreativeTypes: creatives,
		},
	}
	var resp creative.UpdateCreativeResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
