package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/adgroup"
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
