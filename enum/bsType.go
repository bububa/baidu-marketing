package enum

// BsType 该字段仅支持用于Filter中筛选数据使用。
type BsType Enum[int]

func (e BsType) Values() []BsType {
	return []BsType{
		{1, "通报告"},
		{3, "商品报告"},
		{7, "信息流RTA"},
	}
}

func (e *BsType) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = BsType(v)
	return nil
}
