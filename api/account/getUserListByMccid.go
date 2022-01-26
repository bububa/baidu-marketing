package account

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/account"
)

// 账户管家管理
// 账户管家查询下辖子账户，此接口仅供账户管家权限使用，查询操作账户即账户管家的下辖账户列表，header中的被操作账户可以传子账户中任意一个，对该接口功能无影响
func GetUserListByMccid(clt *core.SDKClient, auth model.RequestHeader) ([]account.MccUser, error) {
	req := &model.Request{
		Header: auth,
		Body:   account.GetUserListByMccidRequest{},
	}
	var resp account.GetUserListByMccidResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
