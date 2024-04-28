package oauth

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model/oauth"
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
