package native

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/adgroup/native"
)

// UpdateAdgroup 更新单元
// 更新推广单元
func UpdateAdgroup(clt *core.SDKClient, auth model.RequestHeader, adgroups []native.Adgroup) (*model.ResponseHeader, []native.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body: native.UpdateAdgroupRequest{
			AdgroupFeedTypes: adgroups,
		},
	}
	var resp native.UpdateAdgroupResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
