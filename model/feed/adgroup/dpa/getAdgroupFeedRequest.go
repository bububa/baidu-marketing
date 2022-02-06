package dpa

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetAdgroupFeedRequest 查询推广单元 API Request
type GetAdgroupFeedRequest struct {
	// AdgroupFeedFields 待查询的单元属性
	AdgroupFeedFields []string `json:"adgroupFeedFields,omitempty"`
	// Ids 待查询计划/单元ID集合
	Ids []int64 `json:"ids"`
	// IdType ID类型; 1 - 计划类型; 2 - 单元类型
	IdType int `json:"idType,omitempty"`
}

func (r GetAdgroupFeedRequest) Url() string {
	return fmt.Sprintf("%sDpaAdgroupFeedService/getAdgroupFeed", model.BASE_URL_FEED)
}
