package trans

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/adgroup/trans"
)

// UpdateOcpcTransFeed 修改转化追踪
func UpdateOcpcTransFeed(clt *core.SDKClient, auth *model.RequestHeader, list []trans.OcpcTransFeed) (*model.ResponseHeader, []trans.OcpcTransFeed, error) {
	req := &model.Request{
		Header: auth,
		Body:   &trans.UpdateOcpcTransFeedRequest{TransTraceUpdateTypes: list},
	}
	var resp trans.UpdateOcpcTransFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
