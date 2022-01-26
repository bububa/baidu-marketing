package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// 账户实时数据
// 通过此接口获取账户和计划当天累计数据信息。
func GetAccountLiveData(clt *core.SDKClient, auth model.RequestHeader, dataType int, device int) ([]report.AccountLiveData, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetAccountLiveDataRequest{
			DataType: dataType,
			Device:   device,
		},
	}
	var resp report.GetAccountLiveDataResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
