package enum

// FeedFlowType
type FeedFlowType Enum[int]

func (e FeedFlowType) Values() []FeedFlowType {
	return []FeedFlowType{
		{1, "百度信息流"},
		{2, "贴吧"},
		{3, "好看视频"},
		{4, "百度小说"},
		{5, "百青藤"},
		{6, "默认"},
	}
}

func (e *FeedFlowType) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = FeedFlowType(v)
	return nil
}
