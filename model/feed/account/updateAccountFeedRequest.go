package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type UpdateAccountFeedRequest struct {
	Budget float64 `json:"budget,omitempty"` // 账户预算。默认为0,表示不限预算。正常取值范围为[50-9999999.99]
}

func (r UpdateAccountFeedRequest) Url() string {
	return fmt.Sprintf("%sAccountFeedService/updateAccountFeed", model.BASE_URL_FEED)
}
