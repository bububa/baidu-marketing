package oauth

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetUserInfoRequest 查询授权用户信息
type GetUserInfoRequest struct {
	// OpenID 授权用户查询标识
	OpenID string `json:"openId,omitempty"`
	// AccessToken 已有的授权令牌
	AccessToken string `json:"accessToken,omitempty"`
	// UserID 同意授权的推广账户ID
	UserID uint64 `json:"userId,omitempty"`
	// NeedSubList 是否需要子账号列表，值为true时返回subUserList
	NeedSubList bool `json:"needSubList,omitempty"`
	// PageSize 分页数量，默认100，最大不超过500
	PageSize int `json:"pageSize,omitempty"`
	// LastPageMaxUcId 上一页返回的最大userid，用于子账号列表分页
	// 查询子账号列表时，该字段为必填。第一次获取子账户列表时，该字段需要设置为1
	LastPageMaxUcId uint64 `json:"lastPageMaxUcId,omitempty"`
}

// Url implement RequestBody interface
func (r GetUserInfoRequest) Url() string {
	return util.StringsJoin(model.BASE_OAUTH_URL, "getUserInfo")
}

// UserInfo 授权用户信息
type UserInfo struct {
	// MasterUid 同意授权用户ucid
	MasterUid uint64 `json:"masterUid,omitempty"`
	// MasterName 同意授权用户对应的ucname
	MasterName string `json:"masterName,omitempty"`
	// UserAcctType 授权账户类型
	// 1: 普通账户
	// 2：超管账户（客户中心和账户管家）
	UserAcctType int `json:"userAcctType,omitempty"`
	// HasNext 是否有下一页子账号列表
	HasNext bool `json:"hasNext,omitempty"`
	// SubUserList 同意授权用户关联的子账号列表
	SubUserList []SubUserInfo `json:"subUserList,omitempty"`
}

// SubUserInfo 同意授权用户关联的子账号
type SubUserInfo struct {
	// UcId 同意授权用户关联的子账号ucid
	UcId uint64 `json:"ucId,omitempty"`
	// UcName 同意授权用户关联的子账号ucname
	UcName string `json:"ucName,omitempty"`
}
