package enum

// WinfoIdTypeEnum 定向类型
type WinfoIdType Enum[int]

func (e WinfoIdType) Values() []WinfoIdType {
	return []WinfoIdType{
		{0, "关键词"},
		{1, "词包"},
		{3, "自动扩量"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *WinfoIdType) UnmarshalJSON(b []byte) error {
	val, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = WinfoIdType(val)
	return nil
}
