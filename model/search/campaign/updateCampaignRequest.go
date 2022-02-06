package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// UpdateCampaignRequest 更新计划 API Request
type UpdateCampaignRequest struct {
	// CampaignTypes 更新推广计划字段;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignTypes,omitempty"`
}

func (r UpdateCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/updateCampaign", model.BASE_URL_SMS)
}
