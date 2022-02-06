package campaign

// GetCampaignResponse 查询计划 API Response
type GetCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
