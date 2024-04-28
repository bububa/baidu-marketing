package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetBindBusinessPointListRequest 查询计划已使用的推广业务 API Request
type GetBindBusinessPointListRequest struct {
	// NeedPath 是否需要将完整路径填充到响应的paths字段; true-需要，false-不需要
	NeedPath bool `json:"needPath"`
}

func (r GetBindBusinessPointListRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "BusinessPointService/getBindBusinessPointList")
}

// GetBindBusinessPointListResponse 查询计划已使用的推广业务 API Response
type GetBindBusinessPointListResponse struct {
	Data struct {
		// Paths 推广业务类目完整链路
		Paths []BusinessPoint `json:"paths,omitempty"`
	} `json:"data,omitempty"`
}

type BusinessPoint struct {
	// BusinessPointId 业务类目ID
	BusinessPointId uint64 `json:"businessPointId,omitempty"`
	// BusinessPointName 业务类目名称
	BusinessPointName string `json:"businessPointName,omitempty"`
}
