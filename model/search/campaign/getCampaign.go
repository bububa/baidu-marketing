package campaign

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetCampaignRequest 查询计划 API Request
type GetCampaignRequest struct {
	// CampaignFields 需要查询的计划属性
	// 取值范围：枚举值，列表如下
	// campaignId - 计划ID
	// campaignName - 计划名称
	// budget - 计划每日预算
	// budgetOfflineTime - 预算下线时间
	// exactNegativeWords - 精确否定关键词列表
	// regionTarget - 计划推广地域
	// negativeWords - 短语否定关键词列表
	// pause - 暂停状态
	// schedule - 计划推广暂停时段
	// status - 计划状态
	// marketingTargetId - 营销目标类型
	// adType - 计划类型
	// businessPointId - 推广业务ID
	// businessPointName - 推广业务字面
	// smartRegion - 商品计划: 智能地域开关
	// paDevice - 商品计划: 计划的投放设备
	// os - 商品计划: 计划投放设备平台
	// regionPriceFactor - 分地域出价系数
	// schedulePriceFactors - 分时段出价系数
	// shopType - 电商店铺类型
	// createTime - 添加时间
	// equipmentType - 推广设备
	// campaignBidType - 计划出价方式
	// campaignBid - 计划点击出价
	// campaignOcpcBidType - 计划出价模式
	// campaignOcpcBid - 转化计划出价
	// campaignCvSources - 数据来源
	// campaignTransTypes - 计划目标转化
	// campaignDeepTransTypes - 计划深度转化
	// transAsset - 转化资产类型
	// transAssetId - 转化资产ID
	// geoLocationStatus - 推广地域地理位置选项
	CampaignFields []string `json:"campaignFields,omitempty"`
	// CampaignIds 查询推广计划ID集合
	// 集合长度限制：[0, 100]
	// 输入空返回整个账户的计划ID
	CampaignIds []uint64 `json:"campaignIds,omitempty"`
	// AdType 投放广告类型
	// 取值范围：枚举值，列表如下
	// 0 - 普通计划
	// 14 - 商品计划
	// 不传默认返回全部
	AdType *int `json:"adType,omitempty"`
}

func (r GetCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CampaignService/getCampaign")
}

// GetCampaignResponse 查询计划 API Response
type GetCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
