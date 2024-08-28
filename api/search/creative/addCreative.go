package creative

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// AddCreative 添加基础创意
func AddCreative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *creative.AddCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.AddCreativeResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
