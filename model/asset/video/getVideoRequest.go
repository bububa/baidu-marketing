package video

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetVideoRequest struct {
	Ids      []int `json:"ids,omitempty"` // 视频id集合，不填则不限定视频id。ids大小建议1000以内
	Page     int   `json:"page"`          // 当前页
	PageSize int   `json:"pageSize"`      // 当前页大小
}

func (r GetVideoRequest) Url() string {
	return fmt.Sprintf("%sVideoFeedService/getVideoFeed", model.BASE_URL_FEED)
}
