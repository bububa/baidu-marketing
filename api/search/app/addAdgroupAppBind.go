package app

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/app"
)

// AddAdgroupAppBind 添加APP绑定
func AddAdgroupAppBind(clt *core.SDKClient, auth *model.RequestHeader, reqBody *app.AddAdgroupAppBindRequest) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.AddAdgroupAppBindResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
