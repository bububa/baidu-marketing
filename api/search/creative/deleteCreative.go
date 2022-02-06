package creative

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/creative"
)

// DeleteCreative 删除推广创意
func DeleteCreative(clt *core.SDKClient, auth model.RequestHeader, creativeIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.DeleteCreativeRequest{
			CreativeIds: creativeIds,
		},
	}

	return clt.Do(req, nil)
}
