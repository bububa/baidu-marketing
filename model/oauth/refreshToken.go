package oauth

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// RefreshTokenRequest 更新授权令牌接口
type RefreshTokenRequest struct {
	// AppID 应用ID
	AppID string `json:"appId,omitempty"`
	// RefreshToken 已有的更新令牌 refreshToken
	RefreshToken string `json:"refreshToken,omitempty"`
	// SecretKey 应用密钥
	SecretKey string `json:"secretKey,omitempty"`
	// UserID 同意授权的推广账户ID
	UserID uint64 `json:"userId,omitempty"`
}

// Url implement RequestBody interface
func (r RefreshTokenRequest) Url() string {
	return util.StringsJoin(model.BASE_OAUTH_URL, "refreshToken")
}
