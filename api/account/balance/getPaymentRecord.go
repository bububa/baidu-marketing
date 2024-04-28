package balance

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/account/balance"
)

// GetPaymentRecord 查询付款信息与待加款信息
// 查询付款记录接口，支持获取直销客户实际资金流动、直销客户优惠资金流动、KA实际资金流动、KA优惠资金流动、KA待加款记录。
func GetPaymentRecord(clt *core.SDKClient, auth *model.RequestHeader, reqBody *balance.GetPaymentRecordRequest) (*model.ResponseHeader, *balance.GetPaymentRecordResponse, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp balance.GetPaymentRecordResponse
	header, err := clt.Do(req, &resp)
	if err != nil && header != nil && header.Status != 1 {
		return header, nil, err
	}
	return header, &resp, nil
}
