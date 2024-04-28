package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/creative"
)

// DeleteCreative 删除推广创意
func DeleteCreative(clt *core.SDKClient, auth *model.RequestHeader, creativeFeedIds ...uint64) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: &creative.DeleteCreativeRequest{
			CreativeFeedIds: creativeFeedIds,
		},
	}
	var resp creative.DeleteCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
