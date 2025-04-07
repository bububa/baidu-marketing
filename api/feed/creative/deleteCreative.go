package creative

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/creative"
)

// DeleteCreative 删除推广创意
func DeleteCreative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, creativeFeedIds ...uint64) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: &creative.DeleteCreativeRequest{
			CreativeFeedIds: creativeFeedIds,
		},
	}
	var resp creative.DeleteCreativeResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
