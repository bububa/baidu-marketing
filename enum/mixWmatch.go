package enum

// MixWmatchEnum 匹配模式
type MixWmatch Enum[int]

func (e MixWmatch) Values() []MixWmatch {
	return []MixWmatch{
		{0, "智能匹配"},
		{127, "分匹配出价"},
		{16, "智能匹配核心词"},
		{17, "短语匹配"},
		{48, "精确匹配"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *MixWmatch) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = MixWmatch(v)
	return nil
}
