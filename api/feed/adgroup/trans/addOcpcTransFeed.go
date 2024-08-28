package trans

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/adgroup/trans"
)

// AddOcpcTransFeed 添加转化追踪
func AddOcpcTransFeed(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, list []trans.OcpcTransFeed) (*model.ResponseHeader, []trans.OcpcTransFeed, error) {
	req := &model.Request{
		Header: auth,
		Body:   &trans.AddOcpcTransFeedRequest{OcpcTransFeedTypes: list},
	}
	var resp trans.AddOcpcTransFeedResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
