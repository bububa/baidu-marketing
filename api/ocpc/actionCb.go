package ocpc

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
)

// ActionCb APP转化数据收集
func ActionCb(clt *core.SDKClient, req model.ActionCbRequest) error {
	return clt.ActionCb(req)
}
