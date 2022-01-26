package campaign

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/campaign"
)

// 添加计划
// 新增推广计划，新增时可设置计划的属性设置。
func AddCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.AddCampaignRequest) ([]campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.AddCampaignResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
