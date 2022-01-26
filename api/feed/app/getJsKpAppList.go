package campaign

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/app"
)

// 查询计划
// 根据指定的计划ID获取推广计划(ID可批量)
func GetJsKpAppList(clt *core.SDKClient, auth model.RequestHeader, reqBody *app.GetJsKpAppListRequest) ([]app.ListApp, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetJsKpAppListResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
