package model

/// RegionPriceFactor 地域出价系数
type RegionPriceFactor struct {
	// RegionId 地域ID
	RegionId int `json:"regionId,omitempty"`
	// PriceFactor 出价系数; 取值范围：[1.0, 10.0]
	PriceFactor float64 `json:"priceFactor,omitempty"`
}
