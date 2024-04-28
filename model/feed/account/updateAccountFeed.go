package account

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// UpdateAccountFeedRequest 更新账户信息 API Request
type UpdateAccountFeedRequest struct {
	// AccountFeedType
	AccountFeedType *Account `json:"accountFeedType,omitempty"`
}

func (r UpdateAccountFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AccountFeedService/updateAccountFeed")
}

// UpdateAccountFeedResponse 更新账户信息 API Response
type UpdateAccountFeedResponse struct {
	Data []Account `json:"data,omitempty"`
}
