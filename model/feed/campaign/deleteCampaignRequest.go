package campaign

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// DeleteCampaignRequest 删除计划 API Request
type DeleteCampaignRequest struct {
	// CampaignIds 计划ID
	CampaignIds []int64 `json:"campaignFeedIds,omitempty"`
}

func (r DeleteCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignFeedService/deleteCampaignFeed", model.BASE_URL_FEED)
}
