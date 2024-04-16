package adgroup

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/adgroup"
)

// DeleteAdgroup 删除单元
func DeleteAdgroup(clt *core.SDKClient, auth *model.RequestHeader, adgroupIds ...uint64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.DeleteAdgroupRequest{
			AdgroupIds: adgroupIds,
		},
	}

	return clt.Do(req, nil)
}
