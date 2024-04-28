package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetAdgroupFeedRequest 查询推广单元 API Request
type GetAdgroupFeedRequest struct {
	// AdgroupFeedFields 待查询的单元属性
	// 取值范围：枚举值，列表如下
	// adgroupFeedId - 推广单元ID
	// campaignFeedId - 推广计划ID
	// adgroupFeedName - 推广单元名称
	// pause - 暂停/启用推广单元
	// status - 推广单元状态
	// bid - 出价
	// ftypes - 投放范围
	// bidtype - 优化目标和付费模式
	// ocpc - oCPC信息
	// atpFeedId - 定向包ID
	// addtime - 添加时间
	// modtime - 修改时间
	// deliveryType - 投放场景
	// unefficientAdgroup - 低效单元
	// productSetId - 商品组ID（仅商品推广）
	// ftypeSelection - 是否使用计划流量
	// bidSource - 是否使用计划出价
	// unitOcpxStatus - 单元学习状态
	// atpName - 定向包名称
	// 注意：查询定向设置时，不应传"audience"，而是传原生定向设置列表中需要的字段名称。
	// 某些定向信息内容可能较长（如意图词），请按照实际需要查询，以提高响应速度
	AdgroupFeedFields []string `json:"adgroupFeedFields,omitempty"`
	// Ids 待查询计划/单元ID集合
	// 集合长度限制：[0, 100]
	Ids []uint64 `json:"ids,omitempty"`
	// IdType ID类型; 1 - 计划类型; 2 - 单元类型
	IdType int `json:"idType,omitempty"`
}

func (r GetAdgroupFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "AdgroupFeedService/getAdgroupFeed")
}

// GetAdgroupFeedResponse 查询推广单元 API Response
type GetAdgroupFeedResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
