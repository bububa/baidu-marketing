package enum

// MarketingTarget 营销目标
type MarketingTarget Enum[int]

func (e MarketingTarget) Values() []MarketingTarget {
	return []MarketingTarget{
		{0, "站链接"},
		{1, "应用推广"},
		{2, "门店推广"},
		{3, "推广营销活动"},
		{4, "电商店铺推广"},
		{5, "商品目录"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *MarketingTarget) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = MarketingTarget(v)
	return nil
}
