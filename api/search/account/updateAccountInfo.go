package account

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/account"
)

// 更新账户
// 更新username对应的账户信息，如果是来自MCC Token的请求，则更新target对应的账户信息。
func UpdateAccountInfo(clt *core.SDKClient, auth model.RequestHeader, accountInfo *account.Account) ([]account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.UpdateAccountInfoRequest{
			AccountInfo: accountInfo,
		},
	}
	var resp account.UpdateAccountInfoResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
