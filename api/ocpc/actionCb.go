package ocpc

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
)

// ActionCb APP转化数据收集
func ActionCb(ctx context.Context, clt *core.SDKClient, req model.ActionCbRequest) error {
	return clt.ActionCb(ctx, req)
}
