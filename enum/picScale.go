package enum

// PicScale 图片比例
type PicScale Enum[int]

func (e PicScale) Values() []PicScale {
	return []PicScale{
		{1, "1:1"},
		{2, "1.61:1"},
		{3, "3:1"},
		{4, "1.77:1"},
	}
}

func (e *PicScale) UnmarshalJSON(b []byte) error {
	value, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = PicScale(value)
	return nil
}
