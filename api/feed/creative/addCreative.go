package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/creative"
)

// AddCreative 新增推广创意
func AddCreative(clt *core.SDKClient, auth *model.RequestHeader, reqBody *creative.AddCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.AddCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
