package enum

// ScoreLevel 落地页体验
type ScoreLevel Enum[int]

func (e ScoreLevel) Values() []ScoreLevel {
	return []ScoreLevel{
		{0, "数据积累中"},
		{1, "低于平均"},
		{2, "平均水平"},
		{3, "高于平均"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *ScoreLevel) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = ScoreLevel(val)
	return nil
}
