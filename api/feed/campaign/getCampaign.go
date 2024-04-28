package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/feed/campaign"
)

// GetCampaign 查询计划
// 根据指定的计划ID获取推广计划(ID可批量)
func GetCampaign(clt *core.SDKClient, auth *model.RequestHeader, reqBody *campaign.GetCampaignFeedRequest) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.GetCampaignFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
