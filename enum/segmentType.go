package enum

// SegmentType 比例
type SegmentType Enum[int]

func (e SegmentType) Values() []SegmentType {
	return []SegmentType{
		{101, "1:1"},
		{100, "1:1"},
		{320, "3:1"},
		{321, "1.77:1"},
	}
}

func (e *SegmentType) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = SegmentType(val)
	return nil
}
