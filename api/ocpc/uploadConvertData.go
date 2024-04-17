package ocpc

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/ocpc"
)

// UploadConvertData 广告主回传转化数据接口
// 广告主通过调用该接口，将匹配到的转化数据发送给百度服务器。
func UploadConvertData(clt *core.SDKClient, req *ocpc.UploadConvertDataRequest) (*model.ResponseHeader, error) {
	return clt.Conversion(req, nil)
}
