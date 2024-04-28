package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/adgroup"
)

// GetAdgroup 查询单元
// 查询推广单元
func GetAdgroup(clt *core.SDKClient, auth *model.RequestHeader, reqBody *adgroup.GetAdgroupRequest) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp adgroup.GetAdgroupResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
