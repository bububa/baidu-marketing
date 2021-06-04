package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type UpdateCampaignRequest struct {
	CampaignTypes []Campaign `json:"campaignTypes,omitempty"` // 更新推广计划字段;集合长度限制：[1, 100]
}

func (r UpdateCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/updateCampaign", model.BASE_URL_SMS)
}
