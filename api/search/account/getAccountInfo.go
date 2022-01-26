package account

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/account"
)

// 查询账户
// 获取username对应的账户信息；如果是来自MCC Token的请求，则返回target对应的账户信息。
func GetAccountInfo(clt *core.SDKClient, auth model.RequestHeader, accountFields []string) ([]account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.GetAccountInfoRequest{
			AccountFields: accountFields,
		},
	}
	var resp account.GetAccountInfoResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
