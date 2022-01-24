package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

/*
GetCampaignFeedRequest
CampaignFeedFields 取值范围：枚举值，列表如下
campaignFeedId - 信息流计划Id
campaignFeedName - 信息流计划名称
subject - 推广对象
appinfo - 推广app信息
budget - 推广计划预算
starttime - 推广开始时间
endtime - 推广结束时间
schedule - 推广计划暂停时段
bgtctltype - 预算分配控制方式
pause - 计划启停
status - 推广计划状态
bstype - 物料类型
addtime - 添加时间
shadow - 计划影子的APP信息
*/
type GetCampaignFeedRequest struct {
	CampaignFeedFields []string            `json:"campaignFeedFields,omitempty"` // 需要查询的计划属性
	CampaignFeedIds    []int64             `json:"campaignFeedIds,omitempty"`    // 查询推广计划ID集合
	CampaignFeedFilter *CampaignFeedFilter `json:"campaignFeedFilter,omitempty"` // 计划查询过滤条件
}

func (r GetCampaignFeedRequest) Url() string {
	return fmt.Sprintf("%sCampaignFeedService/getCampaignFeed", model.BASE_URL_FEED)
}

/* CampaignFeedFilter 计划查询过滤条件
bstype 取值范围：枚举值，列表如下
1 - 普通计划
3 - 商品计划
7 - 原生RTA
不填返回全部
*/
type CampaignFeedFilter struct {
	// BsType 计划类型
	BsType []int `json:"bstype,omitempty"`
}
