package campaign

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// AddCampaignRequest 添加计划 API Request
type AddCampaignRequest struct {
	// CampaignTypes 新增推广计划物料;集合长度限制：[1, 100]
	CampaignTypes []Campaign `json:"campaignTypes,omitempty"`
	// AdType 投放广告类型
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 普通计划
	// 14 - 商品计划
	AdType int `json:"adType,omitempty"`
}

func (r AddCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CampaignService/addCampaign")
}

// AddCampaignResponse 添加计划 API Response
type AddCampaignResponse struct {
	Data []Campaign `json:"data,omitempty"`
}
