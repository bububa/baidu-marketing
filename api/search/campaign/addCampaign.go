package campaign

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/campaign"
)

// AddCampaign 添加计划
// 新增推广计划，新增时可设置计划的属性设置。
func AddCampaign(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *campaign.AddCampaignRequest) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.AddCampaignResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
