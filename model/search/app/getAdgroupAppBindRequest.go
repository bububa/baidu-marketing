package app

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetAdgroupAppBindRequest 查询APP绑定
type GetAdgroupAppBindRequest struct {
	// IDType 查询id的层级
	IdType int `json:"idType,omitempty"`
	// Ids id列表
	Ids []int64 `json:"ids,omitempty"`
	// Name app名称
	Name string `json:"name,omitempty"`
	// Platform app操作系统
	Platform []int `json:"platform,omitempty"`
	// Status 状态
	Status []int `json:"status,omitempty"`
	// OrderBy 排序字段, 默认修改时间排序，支持platform,status
	OrderBy string `json:"orderBy,omitempty"`
	// Desc 是否降序
	Desc bool `json:"desc,omitempty"`
	// Limit 分页信息，第一个参数偏移量，第二个参数是页面大小
	Limit []int `json:"limit,omitempty"`
}

func (r GetAdgroupAppBindRequest) Url() string {
	return fmt.Sprintf("%sAdgroupAppService/getAdgroupAppBind", model.BASE_URL_SMS)
}
