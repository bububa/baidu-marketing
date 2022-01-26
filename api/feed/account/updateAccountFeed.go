package account

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/account"
)

// 更新账户信息
func UpdateAccountFeed(clt *core.SDKClient, auth model.RequestHeader, budget float64) ([]account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.UpdateAccountFeedRequest{
			Budget: budget,
		},
	}
	var resp account.UpdateAccountFeedResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
