package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/campaign"
)

// UpdateCampaign 更新计划
// 根据指定的计划ID更新推广计划的属性
func UpdateCampaign(clt *core.SDKClient, auth *model.RequestHeader, campaigns []campaign.Campaign) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body: campaign.UpdateCampaignRequest{
			CampaignTypes: campaigns,
		},
	}
	var resp campaign.UpdateCampaignResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
