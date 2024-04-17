package ocpc

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
)

// ActionCb APP转化数据收集
func ActionCb(clt *core.SDKClient, req model.ActionCbRequest) error {
	return clt.ActionCb(req)
}
