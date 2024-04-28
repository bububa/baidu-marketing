package enum

// QueryStatus 账户添加状态
type QueryStatus Enum[int]

func (e QueryStatus) Values() []QueryStatus {
	return []QueryStatus{
		{0, "已添加"},
		{1, "未添加"},
		{2, "不可添加"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *QueryStatus) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = QueryStatus(val)
	return nil
}
