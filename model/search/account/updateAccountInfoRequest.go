package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type UpdateAccountInfoRequest struct {
	AccountInfo *Account `json:"accountInfo,omitempty"` // 更新账户信息
}

func (r UpdateAccountInfoRequest) Url() string {
	return fmt.Sprintf("%sAccountService/updateAccountInfo", model.BASE_URL_SMS)
}
