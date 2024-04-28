package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/adgroup"
)

// GetAdgroup 查询单元
func GetAdgroup(clt *core.SDKClient, auth *model.RequestHeader, reqBody *adgroup.GetAdgroupFeedRequest) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp adgroup.GetAdgroupFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
