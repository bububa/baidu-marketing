package campaign

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetCampaignFeedRequest 查询计划 API Request
type GetCampaignFeedRequest struct {
	// CampaignFeedFields 需要查询的计划属性
	// 取值范围：枚举值，列表如下
	// campaignFeedId - 信息流计划Id
	// campaignFeedName - 信息流计划名称
	// subject - 营销目标
	// appinfo - 推广app信息
	// budget - 推广计划预算
	// starttime - 推广开始时间
	// endtime - 推广结束时间
	// schedule - 推广计划暂停时段
	// pause - 计划启停
	// status - 推广计划状态
	// bstype - 物料类型
	// campaignType - 计划类型
	// addtime - 添加时间
	// eshopType - 交易所在平台
	// shadow - 计划影子的APP信息
	// budgetOfflineTime - 当天计划预算下线最近一次的时间
	// rtaStatus - 是否开通RTA
	// bid - 出价
	// bidtype - 出价方式
	// ftypes - 投放范围
	// ocpc - oCPC信息
	// unefficientCampaign - 低效计划
	// campaignOcpxStatus - 计划学习状态
	// inheritAscriptionType - 继承归属
	// inheritUserids - 继承优质计划账户id集合
	// inheritCampaignInfos - 继承优质计划的计划信息集合
	// bmcUserId - 商品中心用户ID
	// catalogId - 商品目录ID
	// productType - 产品库类型
	// projectFeedId - 项目ID
	// useLiftBudget - 是否开启一键起量
	// liftBudget - 起量预算
	// liftStatus - 起量状态
	// deliveryType - 投放场景
	// appSubType - 应用推广子类型
	// miniProgramType - 小程序子类型
	// bidMode - 出价模式
	// productIds - 产品ID
	CampaignFeedFields []string `json:"campaignFeedFields,omitempty"`
	// CampaignFeedIds 查询推广计划ID集合
	// 集合长度限制：[0, 100]
	// 输入空返回整个账户的计划ID
	CampaignFeedIds []uint64 `json:"campaignFeedIds,omitempty"`
	// CampaignFeedFilter 计划查询过滤条件
	CampaignFeedFilter *CampaignFeedFilter `json:"campaignFeedFilter,omitempty"`
}

// CampaignFeedFilter 计划查询过滤条件
type CampaignFeedFilter struct {
	// BsType 计划类型
	//     取值范围：枚举值，列表如下
	// 1 - 普通计划
	// 3 - 商品计划
	// 7 - 原生RTA
	// 不填返回全部
	BsType []int `json:"bstype,omitempty"`
}

func (r GetCampaignFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CampaignFeedService/getCampaignFeed")
}

// GetCampaignFeedResponse 查询计划 API Response
type GetCampaignFeedResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
