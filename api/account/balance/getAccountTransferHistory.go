package balance

import (
	"context"
	"time"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account/balance"
)

// GetAccountTransferHistory 查询转账记录
// 查询MCC账户历史转账记录，仅适用于KA账户
func GetAccountTransferHistory(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, startTime time.Time, endTime time.Time) (*model.ResponseHeader, []balance.AccountTransferHistory, error) {
	req := &model.Request{
		Header: auth,
		Body: &balance.GetAccountTransferHistoryRequest{
			StartTime: startTime.Format("2006-01-02"),
			EndTime:   endTime.Format("2006-01-02"),
		},
	}
	var resp balance.GetAccountTransferHistoryResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
