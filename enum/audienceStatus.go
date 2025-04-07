package enum

// AudienceStatusEnum 类型
type AudienceStatus Enum[int]

func (e AudienceStatus) Values() []AudienceStatus {
	return []AudienceStatus{
		{0, "基本属性"},
		{1, "自定义人群"},
		{1025, "已转化用户"},
		{2, "设备属性"},
		{4, "商圈地域"},
	}
}

func (e *AudienceStatus) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = AudienceStatus(v)
	return nil
}
