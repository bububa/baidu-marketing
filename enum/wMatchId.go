package enum

// WMatchId 触发模式
type WMatchId Enum[int]

func (e WMatchId) Values() []WMatchId {
	return []WMatchId{
		{101, "自动定向（智能匹配）"},
		{103, "网址投放"},
		{109, "商品定向"},
		{110, "自动定向"},
		{111, "词包"},
		{15, "智能"},
		{16, "智能匹配-人群智选"},
		{31, "短语"},
		{63, "精确"},
	}
}

func (e *WMatchId) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = WMatchId(val)
	return nil
}
