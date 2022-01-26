package app

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/app"
)

func GetAppList(clt *core.SDKClient, auth model.RequestHeader, reqBody *app.GetAppListRequest) ([]app.InfoApp, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAppListResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
