package account

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetAccountFeedRequest struct {
	AccountFeedFields []string `json:"accountFeedFields,omitempty"` // 需要查询的账户属性。取值范围参考账户信息对象中的字段取值
}

func (r GetAccountFeedRequest) Url() string {
	return fmt.Sprintf("%sAccountFeedService/getAccountFeed", model.BASE_URL_FEED)
}
