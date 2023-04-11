package campaign

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// AddCampaignRequest 添加计划 API Request
type AddCampaignRequest struct {
	// CampaignTypes 新增推广计划物料;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignFeedTypes,omitempty"`
}

func (r AddCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignFeedService/addCampaignFeed", model.BASE_URL_FEED)
}
