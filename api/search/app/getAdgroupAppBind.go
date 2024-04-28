package app

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/app"
)

// GetAdgroupAppBind 查询APP绑定
func GetAdgroupAppBind(clt *core.SDKClient, auth *model.RequestHeader, reqBody *app.GetAdgroupAppBindRequest) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAdgroupAppBindResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
