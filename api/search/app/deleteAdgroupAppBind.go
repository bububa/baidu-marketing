package app

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/app"
)

// DeleteAdgroupAppBind 删除APP绑定
func DeleteAdgroupAppBind(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, bindIds ...uint64) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   &app.DeleteAdgroupAppBindRequest{DelBindIds: bindIds},
	}
	var resp app.DeleteAdgroupAppBindResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
