package campaign

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// DeleteCampaignRequest 删除计划 API Request
type DeleteCampaignRequest struct {
	// CampaignIds 计划ID
	CampaignIds []uint64 `json:"campaignFeedIds,omitempty"`
}

func (r DeleteCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CampaignFeedService/deleteCampaignFeed")
}

// DeleteCampaignResponse 删除计划 API Response
type DeleteCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
