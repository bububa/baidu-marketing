package balance

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account/balance"
)

// 查询待加款信息
// 支持KA账户查询付款记录，此接口仅供KA账户使用
func GetPaymentHistory(clt *core.SDKClient, auth model.RequestHeader, reqBody balance.GetPaymentHistoryRequest) ([]balance.PaymentHistory, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp balance.GetPaymentHistoryResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
