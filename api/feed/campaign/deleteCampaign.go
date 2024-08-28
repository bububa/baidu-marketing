package campaign

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/campaign"
)

// DeleteCampaign 删除计划
func DeleteCampaign(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, campaignIds ...uint64) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body: &campaign.DeleteCampaignRequest{
			CampaignIds: campaignIds,
		},
	}

	var ret campaign.DeleteCampaignResponse
	headers, err := clt.Do(ctx, req, &ret)
	if err != nil {
		return nil, nil, err
	}
	return headers, ret.Data, nil
}
