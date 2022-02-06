package ocpc

import (
	"encoding/json"
	"net/http"

	"github.com/bububa/baidu-marketing/model/ocpc"
)

// ActionCb APP转化数据收集
func ActionCb(req *ocpc.ActionCbRequest) error {
	httpReq, err := http.NewRequest("GET", req.Url(), nil)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	var resp ocpc.ActionCbResponse
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
