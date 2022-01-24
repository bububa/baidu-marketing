package app

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

type GetJsKpAppListRequest struct {
}

func (r GetJsKpAppListRequest) Url() string {
	return fmt.Sprintf("%sAppFeedService/getJsKpAppList", model.BASE_URL_FEED)
}
