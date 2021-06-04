package model

type RegionPriceFactor struct {
	RegionId    int     `json:"regionId,omitempty"`    // 地域ID
	PriceFactor float64 `json:"priceFactor,omitempty"` // 出价系数; 取值范围：[1.0, 10.0]
}
