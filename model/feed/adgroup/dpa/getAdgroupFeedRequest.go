package dpa

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetAdgroupFeedRequest struct {
	AdgroupFeedFields []string `json:"adgroupFeedFields,omitempty"` // 待查询的单元属性
	Ids               []int64  `json:"ids"`                         // 待查询计划/单元ID集合
	IdType            int      `json:"idType,omitempty"`            // ID类型; 1 - 计划类型; 2 - 单元类型
}

func (r GetAdgroupFeedRequest) Url() string {
	return fmt.Sprintf("%sDpaAdgroupFeedService/getAdgroupFeed", model.BASE_URL_FEED)
}
