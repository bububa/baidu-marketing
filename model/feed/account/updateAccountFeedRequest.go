package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// UpdateAccountFeedRequest 更新账户信息 API Request
type UpdateAccountFeedRequest struct {
	// Budget 账户预算。默认为0,表示不限预算。正常取值范围为[50-9999999.99]
	Budget float64 `json:"budget,omitempty"`
}

func (r UpdateAccountFeedRequest) Url() string {
	return fmt.Sprintf("%sAccountFeedService/updateAccountFeed", model.BASE_URL_FEED)
}
