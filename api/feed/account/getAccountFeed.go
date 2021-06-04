package account

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/account"
)

// 查询账户
func GetAccountFeed(clt *core.SDKClient, auth model.RequestHeader, accountFields []string) ([]account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.GetAccountFeedRequest{
			AccountFeedFields: accountFields,
		},
	}
	var resp account.GetAccountFeedResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
