package adgroup

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetAdgroupRequest 查询推广单元API Request
type GetAdgroupRequest struct {
	// Ids 查询id集合;idType=5时，类型为单元ID，不超过5000个；idType=3时，类型为计划ID，不超过100个
	Ids []uint64 `json:"ids,omitempty"`
	// AdgroupFields 查询推广单元字段
	// 取值范围：枚举值，列表如下
	// adgroupId - 推广单元ID
	// campaignId - 推广计划ID
	// adgroupName - 单元名称
	// pause - 推广单元启用/暂停
	// appShopDirectStatus - 应用商店直投
	// maxPrice - 单元出价
	// negativeWords - 单元短语否定关键词
	// exactNegativeWords - 单元精确否定关键词
	// status - 单元状态
	// segmentRecommendStatus - 自动配图开关
	// creativeTextOptimizationStatus - 自动文案优化
	// offlineReasons - 下线理由
	// paPrice - 推广单元商品出价
	// adType - 广告类型
	// monitorUrl - 单元层级监控
	// pcFinalUrl - 计算机最终访问网址
	// pcTrackParam - 计算机监控后缀
	// pcTrackTemplate - 计算机第三方追踪模板
	// mobileFinalUrl - 移动最终访问网址
	// mobileTrackParam - 移动监控后缀
	// mobileTrackTemplate - 移动第三方追踪模板url
	// productSetId -虚拟商品组id
	// createTime-添加时间
	// adgroupAutoTargetingStatus-自动定向
	AdgroupFields []string `json:"adgroupFields,omitempty"`
	// IdType 查询层级;3 - 计划ID;5 - 单元ID
	IdType int `json:"idType,omitempty"`
	// GetTemp 是否查询单元影子
	// 默认值：0
	// 0 - 只查询单元本身，1 - 只查询单元影子；
	// 想要获得单元的全集，需要调用该方法两次，分别为getTemp=0和getTemp=1；
	// 影子说明：用户先向系统提交了单元A，并且A已审核通过，之后再对A进行影响审核状态的修改（修改pcFinalUrl、pcTrackParam、pcTrackTemplate、mobileFinalUrl、mobileTrackParam、mobileTrackTemplate等字段），修改后的单元为A’（A’即为影子，仅对审核通过的物料进行修改才会产生影子），在A’通过审核生效之前，线上的生效创意仍然为A。 此时：getTemp为0查询到的是A，getTemp为1查询到的是A’
	GetTemp int `json:"getTemp,omitempty"`
}

func (r GetAdgroupRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "AdgroupService/getAdgroup")
}

// GetAdgroupResponse 查询推广单元API Response
type GetAdgroupResponse struct {
	Data []Adgroup `json:"data,omitempty"`
}
