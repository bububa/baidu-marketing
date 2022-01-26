package adgroup

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/adgroup"
)

// 更新单元
// 更新推广单元
func UpdateAdgroup(clt *core.SDKClient, auth model.RequestHeader, adgroups []adgroup.Adgroup) ([]adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.UpdateAdgroupRequest{
			AdgroupTypes: adgroups,
		},
	}
	var resp adgroup.UpdateAdgroupResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
