package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// UpdateAccountInfoRequest 更新账户 API Request
type UpdateAccountInfoRequest struct {
	// AccountInfo 更新账户信息
	AccountInfo *Account `json:"accountInfo,omitempty"`
}

func (r UpdateAccountInfoRequest) Url() string {
	return fmt.Sprintf("%sAccountService/updateAccountInfo", model.BASE_URL_SMS)
}
