package balance

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account/balance"
)

// 查询付款信息与待加款信息
// 查询付款记录接口，支持获取直销客户实际资金流动、直销客户优惠资金流动、KA实际资金流动、KA优惠资金流动、KA待加款记录。
func GetPaymentRecord(clt *core.SDKClient, auth model.RequestHeader, reqBody balance.GetPaymentRecordRequest) (*balance.GetPaymentRecordResponse, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp balance.GetPaymentRecordResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}
