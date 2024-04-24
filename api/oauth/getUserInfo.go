package oauth

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model/oauth"
)

// GetUserInfo 查询授权用户信息
func GetUserInfo(clt *core.SDKClient, req *oauth.GetUserInfoRequest) (*oauth.UserInfo, error) {
	var resp oauth.UserInfo
	if err := clt.OAuth(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
