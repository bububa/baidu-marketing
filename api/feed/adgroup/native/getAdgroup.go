package native

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/adgroup/native"
)

// GetAdgroup 查询原生推广单元
func GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *native.GetAdgroupFeedRequest) (*model.ResponseHeader, []native.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp native.GetAdgroupFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
