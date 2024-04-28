package balance

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/account/balance"
)

// GetBalanceInfo 查询账户余额成分
// 支持KA账户&直销客户查询账户余额及余额成分。不同渠道账户使用的资金包不同，详见文档说明
func GetBalanceInfo(clt *core.SDKClient, auth *model.RequestHeader, productIds []int) (*model.ResponseHeader, []balance.BalanceInfo, error) {
	req := &model.Request{
		Header: auth,
		Body: &balance.GetBalanceInfoRequest{
			ProductIds: productIds,
		},
	}
	var resp balance.GetBalanceInfoResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
