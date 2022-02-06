package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetAdgroupRequest 查询推广单元API Request
// 查询推广单元字段
// adgroupId - 推广单元ID
// campaignId - 推广计划ID
// adgroupName - 单元名称
// pause - 推广单元启用/暂停
// maxPrice - 单元出价
// negativeWords - 单元短语否定关键词
// exactNegativeWords - 单元精确否定关键词
// status - 单元状态
// priceRatio - 单元移动出价系数（默认使用计划的）
// pcPriceRatio - 单元计算机出价系数（默认使用计划的）
// segmentRecommendStatus - 基础创意智能配图开关
// offlineReasons - 下线理由
// paPrice - 推广单元商品出价
// adType - 广告类型
// monitorUrl - 单元层级监控url
// productSetId - 虚拟商品组id
type GetAdgroupRequest struct {
	// Ids 查询id集合;idType=5时，类型为单元ID，不超过5000个；idType=3时，类型为计划ID，不超过100个
	Ids []int64 `json:"ids"`
	// AdgroupFields 查询推广单元字段
	AdgroupFields []string `json:"adgroupFields"`
	// IdType 查询层级;3 - 计划ID;5 - 单元ID
	IdType int `json:"idType"`
}

func (r GetAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/getAdgroup", model.BASE_URL_SMS)
}
