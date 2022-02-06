package adgroup

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/adgroup"
)

// AddAdgroup 添加单元
// 新增推广单元
func AddAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *adgroup.AddAdgroupRequest) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp adgroup.AddAdgroupResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
