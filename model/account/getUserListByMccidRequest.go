package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetUserListByMccidRequest 账户管家子账号 API Request
type GetUserListByMccidRequest struct{}

func (r GetUserListByMccidRequest) Url() string {
	return fmt.Sprintf("%sMccFeedService/getUserListByMccid", model.BASE_URL_FEED)
}
