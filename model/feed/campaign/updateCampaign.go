package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// UpdateCampaignRequest 更新计划 API Request
type UpdateCampaignRequest struct {
	// CampaignTypes 更新推广计划字段;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignFeedTypes,omitempty"`
}

func (r UpdateCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CampaignFeedService/updateCampaignFeed")
}

// UpdateCampaignResponse 更新计划 API Response
type UpdateCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
