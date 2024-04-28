package account

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/account"
)

// GetUserListByMccid 账户管家管理
// 账户管家查询下辖子账户，此接口仅供账户管家权限使用，查询操作账户即账户管家的下辖账户列表，header中的被操作账户可以传子账户中任意一个，对该接口功能无影响
func GetUserListByMccid(clt *core.SDKClient, auth *model.RequestHeader) (*model.ResponseHeader, []account.MccUser, error) {
	req := &model.Request{
		Header: auth,
		Body:   new(account.GetUserListByMccidRequest),
	}
	var resp account.GetUserListByMccidResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
