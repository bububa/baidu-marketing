package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetCampaignRequest 查询计划 API Request
type GetCampaignRequest struct {
	// CampaignFields 需要查询的计划属性
	CampaignFields []string `json:"campaignFields,omitempty"`
	// CampaignIds 查询推广计划ID集合
	CampaignIds []int64 `json:"campaignIds,omitempty"`
	// MobileExtend 是否获取移动优先计划
	MobileExtend int `json:"mobileExtend,omitempty"`
	// AdType 投放广告类型;0 - 普通计划;14 - 商品计划;6 - 网址定向计划
	AdType int `json:"adType,omitempty"`
}

func (r GetCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/getCampaign", model.BASE_URL_SMS)
}
