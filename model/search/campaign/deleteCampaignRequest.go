package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type DeleteCampaignRequest struct {
	CampaignIds []int64 `json:"campaignIds,omitempty"`
}

func (r DeleteCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/deleteCampaign", model.BASE_URL_SMS)
}
