package app

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/app"
)

// GetAdgroupAppBind 查询APP绑定
func GetAdgroupAppBind(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *app.GetAdgroupAppBindRequest) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAdgroupAppBindResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
