package trans

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/adgroup/trans"
)

// GetOcpcTransFeed 查询转化追踪
func GetOcpcTransFeed(clt *core.SDKClient, auth *model.RequestHeader, reqBody *trans.GetOcpcTransFeedRequest) (*model.ResponseHeader, []trans.OcpcTransFeed, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp trans.GetOcpcTransFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
