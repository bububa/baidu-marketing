package oauth

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model/oauth"
)

// AccessToken 换取授权令牌接口
func AccessToken(clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.AccessToken, error) {
	var resp oauth.AccessToken
	req.AppID = clt.AppID()
	req.SecretKey = clt.Secret()
	if err := clt.OAuth(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
