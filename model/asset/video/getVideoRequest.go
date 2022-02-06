package video

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetVideoRequest 获取视频素材 API Request
type GetVideoRequest struct {
	// Ids 视频id集合，不填则不限定视频id。ids大小建议1000以内
	Ids []int `json:"ids,omitempty"`
	// Page 当前页
	Page int `json:"page"`
	// PageSize 当前页大小
	PageSize int `json:"pageSize"`
}

func (r GetVideoRequest) Url() string {
	return fmt.Sprintf("%sVideoFeedService/getVideoFeed", model.BASE_URL_FEED)
}
