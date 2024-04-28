package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/creative"
)

// GetCreative 查询基础创意
func GetCreative(clt *core.SDKClient, auth *model.RequestHeader, reqBody *creative.GetCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.GetCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
