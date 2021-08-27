package balance

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account/balance"
)

// 查询账户余额成分
// 支持KA账户&直销客户查询账户余额及余额成分。不同渠道账户使用的资金包不同，详见文档说明
func GetBalanceInfo(clt *core.SDKClient, auth model.RequestHeader, productIds []int) ([]balance.BalanceInfo, error) {
	req := &model.Request{
		Header: auth,
		Body: balance.GetBalanceInfoRequest{
			ProductIds: productIds,
		},
	}
	var resp balance.GetBalanceInfoResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
