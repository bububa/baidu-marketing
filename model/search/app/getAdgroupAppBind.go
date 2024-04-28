package app

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetAdgroupAppBindRequest 查询APP绑定
type GetAdgroupAppBindRequest struct {
	// IDType 查询id的层级
	//     取值范围：枚举值，列表如下
	// 2 - 用户层级
	// 3 - 计划层级
	// 5 - 单元层级
	// 8 - BIND层级
	IdType int `json:"idType,omitempty"`
	// Ids id列表
	Ids []uint64 `json:"ids,omitempty"`
	// Name app名称
	Name string `json:"name,omitempty"`
	// Platform app操作系统
	// 取值范围：枚举值，列表如下
	// 1 - 安卓
	// 3 - iOS
	Platform []int `json:"platform,omitempty"`
	// Status 状态
	// 第一位表示暂停位（0启动，1暂停）；第二位表示APP有效位（0有效，1无效）；第三位表示审核中（0非审核中，1审核中）
	Status []int `json:"status,omitempty"`
	// OrderBy 排序字段, 默认修改时间排序，支持platform,status
	OrderBy string `json:"orderBy,omitempty"`
	// Desc 是否降序
	Desc bool `json:"desc,omitempty"`
	// Limit 分页信息，第一个参数偏移量，第二个参数是页面大小
	Limit []int `json:"limit,omitempty"`
}

func (r GetAdgroupAppBindRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupAppService/getAdgroupAppBind")
}

// GetAdgroupAppBindResponse 查询APP绑定 API Response
type GetAdgroupAppBindResponse struct {
	Data []AppBindItem `json:"data,omitempty"`
}
