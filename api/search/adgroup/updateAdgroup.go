package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/adgroup"
)

// UpdateAdgroup 更新单元
// 更新推广单元
func UpdateAdgroup(clt *core.SDKClient, auth *model.RequestHeader, adgroups []adgroup.Adgroup) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.UpdateAdgroupRequest{
			AdgroupTypes: adgroups,
		},
	}
	var resp adgroup.UpdateAdgroupResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
