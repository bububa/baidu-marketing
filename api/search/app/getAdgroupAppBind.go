package app

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/app"
)

func GetAdgroupAppBind(clt *core.SDKClient, auth model.RequestHeader, reqBody *app.GetAdgroupAppBindRequest) ([]app.InfoAppItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAdgroupAppBindResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
