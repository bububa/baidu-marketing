package app

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/app"
)

// GetAppList 获取APP素材
func GetAppList(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *app.GetAppListRequest) (*model.ResponseHeader, []app.AppInfoItemList, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAppListResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
