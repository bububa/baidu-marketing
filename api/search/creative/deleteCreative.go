package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/creative"
)

// DeleteCreative 删除基础创意
func DeleteCreative(clt *core.SDKClient, auth *model.RequestHeader, creativeIds ...uint64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: &creative.DeleteCreativeRequest{
			CreativeIds: creativeIds,
		},
	}

	return clt.Do(req, nil)
}
