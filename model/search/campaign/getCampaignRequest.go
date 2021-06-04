package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetCampaignRequest struct {
	CampaignFields []string `json:"campaignFields,omitempty"` // 需要查询的计划属性
	CampaignIds    []int64  `json:"campaignIds"`              // 查询推广计划ID集合
	MobileExtend   int      `json:"mobileExtend,omitempty"`   // 是否获取移动优先计划
	AdType         int      `json:"adType,omitempty"`         // 投放广告类型;0 - 普通计划;14 - 商品计划;6 - 网址定向计划
}

func (r GetCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/getCampaign", model.BASE_URL_SMS)
}
