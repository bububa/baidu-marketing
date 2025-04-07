package account

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/account"
)

// GetAccountFeed 查询账户
func GetAccountFeed(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, accountFields []string) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: &account.GetAccountFeedRequest{
			AccountFeedFields: accountFields,
		},
	}
	var resp account.GetAccountFeedResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
