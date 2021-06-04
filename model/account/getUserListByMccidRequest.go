package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetUserListByMccidRequest struct{}

func (r GetUserListByMccidRequest) Url() string {
	return fmt.Sprintf("%sMccFeedService/getUserListByMccid", model.BASE_URL_FEED)
}
