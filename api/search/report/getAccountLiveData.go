package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// GetAccountLiveData 账户实时数据
// 通过此接口获取账户和计划当天累计数据信息。
func GetAccountLiveData(clt *core.SDKClient, auth *model.RequestHeader, dataType int, device int) (*model.ResponseHeader, []report.AccountLiveData, error) {
	req := &model.Request{
		Header: auth,
		Body: &report.GetAccountLiveDataRequest{
			DataType: dataType,
			Device:   device,
		},
	}
	var resp report.GetAccountLiveDataResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
