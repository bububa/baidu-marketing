package oauth

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AccessToken 授权令牌
type AccessToken struct {
	// AccessToken 授权令牌
	AccessToken string `json:"accessToken,omitempty"`
	// RefreshToken 刷新令牌
	RefreshToken string `json:"refreshToken,omitempty"`
	// OpenID 获取授权用户信息标识
	OpenID string `json:"openId,omitempty"`
	// ExpiresTime 授权令牌到期时间
	ExpiresTime string `json:"expiresTime,omitempty"`
	// RefreshExpiresTime 更新令牌到期时间
	RefreshExpiresTime string `json:"refreshExpiresTime,omitempty"`
	// ExpiresIn 授权令牌剩余有效时间
	ExpiresIn int64 `json:"expiresIn,omitempty"`
	// RefreshExpiresIn 更新令牌剩余有效时间
	RefreshExpiresIn int64 `json:"refreshExpiresIn,omitempty"`
	// UserID 授权账号 ucid
	UserID uint64 `json:"userId,omitempty"`
}

// AccessTokenRequest 换取授权令牌接口
type AccessTokenRequest struct {
	// AppID 应用 appid
	AppID string `json:"appId,omitempty"`
	// AuthCode 临时授权码（数据来源：通过回调接口获取）
	AuthCode string `json:"authCode,omitempty"`
	// SecretKey 应用持有的 secretKey
	SecretKey string `json:"secretKey,omitempty"`
	// GrantType 授权令牌获取模式，仅限：auth_code（授权码模式）
	GrantType string `json:"grantType,omitempty"`
	// UserID 同意授权的推广账户ID（数据来源：通过回调接口获取）
	UserID uint64 `json:"userId,omitempty"`
}

// Url implement RequestBody interface
func (r AccessTokenRequest) Url() string {
	return util.StringsJoin(model.BASE_OAUTH_URL, "accessToken")
}
