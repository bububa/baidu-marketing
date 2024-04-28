package account

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetUserListByMccidRequest 账户管家子账号 API Request
type GetUserListByMccidRequest struct{}

// Url implement RequestBody interface
func (r GetUserListByMccidRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "MccFeedService/getUserListByMccid")
}

// GetUserListByMccidResponse 账户管家子账号列表 API Response
type GetUserListByMccidResponse struct {
	Data []MccUser `json:"data,omitempty"`
}

// MccUser 账户管家下辖子账号
type MccUser struct {
	// UserId 用户id
	Userid uint64 `json:"userid,omitempty"`
	// Username 用户名
	Username string `json:"username,omitempty"`
	// Mccid 账户管家id
	Mccid uint64 `json:"mccid,omitempty"`
	// Fatname 账户管家名称
	Fatname string `json:"fatname,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
}
