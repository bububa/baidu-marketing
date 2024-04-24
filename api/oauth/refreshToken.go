package oauth

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model/oauth"
)

// RefreshToken 更新授权令牌接口
func RefreshToken(clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.AccessToken, error) {
	var resp oauth.AccessToken
	req.AppID = clt.AppID()
	req.SecretKey = clt.Secret()
	if err := clt.OAuth(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
