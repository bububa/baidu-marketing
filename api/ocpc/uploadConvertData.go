package ocpc

import (
	"encoding/json"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model/ocpc"
)

// 广告主回传转化数据接口
// 广告主通过调用该接口，将匹配到的转化数据发送给百度服务器。
func UploadConvertData(clt *core.SDKClient, req *ocpc.UploadConvertDataRequest) error {
	if req.Token == "" {
		req.Token = clt.Token()
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var resp ocpc.Response
	err = clt.Post(UPLOAD_CONVERT_DATA_URL, reqBytes, &resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
