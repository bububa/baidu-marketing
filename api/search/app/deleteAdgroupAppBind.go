package app

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/app"
)

// DeleteAdgroupAppBind 删除APP绑定
func DeleteAdgroupAppBind(clt *core.SDKClient, auth *model.RequestHeader, bindIds ...uint64) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   &app.DeleteAdgroupAppBindRequest{DelBindIds: bindIds},
	}
	var resp app.DeleteAdgroupAppBindResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
