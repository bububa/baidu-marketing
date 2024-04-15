package app

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetAppListRequest 获取APP素材 API Request
type GetAppListRequest struct {
	// Platforms app操作系统
	// 取值范围：枚举值，列表如下
	// 1 - 安卓
	// 3 - iOS
	Platforms []int `json:"platforms,omitempty"`
	// Limit 分页信息，第一个参数偏移量，第二个参数是页面大小
	Limit []int `json:"limit,omitempty"`
}

func (r GetAppListRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AppProcessService/getAppList")
}

// GetAppListResponse 获取APP素材 API Response
type GetAppListResponse struct {
	Data []AppInfoItemList `json:"data,omitempty"`
}

type AppInfoItemList struct {
	AppInfoList []AppInfoItem `json:"appInfoList,omitempty"`
	TotalCount  int           `json:"totalCount,omitempty"`
}
