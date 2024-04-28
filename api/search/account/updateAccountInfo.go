package account

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/account"
)

// UpdateAccountInfo 更新账户
// 更新username对应的账户信息，如果是来自MCC Token的请求，则更新target对应的账户信息。
func UpdateAccountInfo(clt *core.SDKClient, auth *model.RequestHeader, accountInfo *account.UpdateAccountInfo) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: &account.UpdateAccountInfoRequest{
			AccountInfo: accountInfo,
		},
	}
	var resp account.UpdateAccountInfoResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
