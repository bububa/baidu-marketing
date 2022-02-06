package adgroup

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/adgroup"
)

// GetAdgroup 查询单元
// 查询推广单元
func GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *adgroup.GetAdgroupRequest) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp adgroup.GetAdgroupResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
