package native

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// AddAdgroupRequest 添加单元 API Request
type AddAdgroupRequest struct {
	// AdgroupFeedTypes 新增推广单元物料
	AdgroupFeedTypes []Adgroup `json:"adgroupFeedTypes"`
}

func (r AddAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupFeedService/addAdgroupFeed", model.BASE_URL_FEED)
}
