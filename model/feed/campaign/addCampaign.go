package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AddCampaignRequest 添加计划 API Request
type AddCampaignRequest struct {
	// CampaignTypes 新增推广计划物料;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignFeedTypes,omitempty"`
}

func (r AddCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CampaignFeedService/addCampaignFeed")
}

// AddCampaignResponse 添加计划 API Response
type AddCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
