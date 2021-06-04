package campaign

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/campaign"
)

// 删除计划
func DeleteCampaign(clt *core.SDKClient, auth model.RequestHeader, campaignIds []int64) error {
	req := &model.Request{
		Header: auth,
		Body: campaign.DeleteCampaignRequest{
			CampaignIds: campaignIds,
		},
	}

	return clt.Do(req, nil)
}
