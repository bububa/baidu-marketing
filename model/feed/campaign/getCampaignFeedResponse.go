package campaign

// GetCampaignFeedResponse 查询计划 API Response
type GetCampaignFeedResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
