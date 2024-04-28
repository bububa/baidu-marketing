package account

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/account"
)

// GetAccountInfo 查询账户
// 获取username对应的账户信息；如果是来自MCC Token的请求，则返回target对应的账户信息。
func GetAccountInfo(clt *core.SDKClient, auth *model.RequestHeader, accountFields []string) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: &account.GetAccountInfoRequest{
			AccountFields: accountFields,
		},
	}
	var resp account.GetAccountInfoResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
