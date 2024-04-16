package campaign

// InheritCampaignInfo 继承计划信息集合
type InheritCampaignInfo struct {
	// UserId 账户ID
	UserId uint64 `json:"userId,omitempty"`
	// Campaigns 继承的计划信息集合
	Campaigns []CampaignIdName `json:"campaigns,omitempty"`
}

type CampaignIdName struct {
	// CampaignId 计划ID
	CampaignId uint64 `json:"campaignId,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaignName,omitempty"`
}
