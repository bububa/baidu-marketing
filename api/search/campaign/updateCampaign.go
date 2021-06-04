package campaign

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/campaign"
)

// 更新计划
// 根据指定的计划ID更新推广计划的属性
func UpdateCampaign(clt *core.SDKClient, auth model.RequestHeader, campaigns []campaign.Campaign) ([]campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body: campaign.UpdateCampaignRequest{
			CampaignTypes: campaigns,
		},
	}
	var resp campaign.UpdateCampaignResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
