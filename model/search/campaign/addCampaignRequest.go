package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// AddCampaignRequest 添加计划 API Request
type AddCampaignRequest struct {
	// CampaignTypes 新增推广计划物料;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignTypes,omitempty"`
	// AdType 投放广告类型;0 - 普通计划;14 - 商品计划;6 - 网址定向计划
	AdType int `json:"adType,omitempty"`
}

func (r AddCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/addCampaign", model.BASE_URL_SMS)
}
