package dpa

import "github.com/bububa/baidu-marketing/model/feed/adgroup"

type Adgroup struct {
	// AdgroupId 推广单元ID，定义同原生
	AdgroupId int64 `json:"adgroupFeedId,omitempty"`
	// CampaignId 所属推广计划ID，定义同原生
	CampaignId int64 `json:"campaignFeedId,omitempty"`
	// AdgroupName 推广单元名称，定义同原生
	AdgroupName string `json:"adgroupFeedName,omitempty"`
	// Ftypes 投放流量，定义同原生
	FTypes []int `json:"ftypes,omitempty"`
	// ProductTypes 投放版位，定义同原生
	ProductTypes []int `json:"producttypes,omitempty"`
	// Pause 是否暂停推广，定义同原生
	Pause *bool `json:"pause,omitempty"`
	// Status 推广单元状态，定义同原生
	Status int `json:"status,omitempty"`
	// ProductSetId 商品商品组ID。参照信息流商品组管理
	ProductSetId int64 `json:"productSetId,omitempty"`
	// Audience 商品单元定向设置
	Audience *Audience `json:"audience,omitempty"`
	// BidType 优化目标、付费模式
	BidType int `json:"bidtype,omitempty"`
	// Ocpc 商品oCPC设置，定义同原生
	Ocpc *adgroup.Ocpc `json:"ocpc,omitempty"`
}
