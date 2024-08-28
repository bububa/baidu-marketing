package adgroup

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/adgroup"
)

// UpdateAdgroup 更新单元
// 更新推广单元
func UpdateAdgroup(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, adgroups []adgroup.Adgroup) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.UpdateAdgroupRequest{
			AdgroupTypes: adgroups,
		},
	}
	var resp adgroup.UpdateAdgroupResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
