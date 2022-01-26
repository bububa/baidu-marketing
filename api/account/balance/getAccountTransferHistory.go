package balance

import (
	"time"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account/balance"
)

// 查询转账记录
// 查询MCC账户历史转账记录，仅适用于KA账户
func GetAccountTransferHistory(clt *core.SDKClient, auth model.RequestHeader, startTime time.Time, endTime time.Time) ([]balance.AccountTransferHistory, error) {
	req := &model.Request{
		Header: auth,
		Body: balance.GetAccountTransferHistoryRequest{
			StartTime: startTime.Format("2006-01-02"),
			EndTime:   endTime.Format("2006-01-02"),
		},
	}
	var resp balance.GetAccountTransferHistoryResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
