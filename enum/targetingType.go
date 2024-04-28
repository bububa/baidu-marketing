package enum

// TargetingType 购买方式
type TargetingType Enum[int]

func (e TargetingType) Values() []TargetingType {
	return []TargetingType{
		{0, "关键词"},
		{6, "网址定向"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *TargetingType) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = TargetingType(val)
	return nil
}
