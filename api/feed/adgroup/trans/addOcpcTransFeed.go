package trans

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/adgroup/trans"
)

// AddOcpcTransFeed 添加转化追踪
func AddOcpcTransFeed(clt *core.SDKClient, auth *model.RequestHeader, list []trans.OcpcTransFeed) (*model.ResponseHeader, []trans.OcpcTransFeed, error) {
	req := &model.Request{
		Header: auth,
		Body:   &trans.AddOcpcTransFeedRequest{OcpcTransFeedTypes: list},
	}
	var resp trans.AddOcpcTransFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
