package enum

// FeedSubjectEnum 该字段仅支持用于Filter中筛选数据使用。
type FeedSubject Enum[int]

func (e FeedSubject) Values() []FeedSubject {
	return []FeedSubject{
		{0, "全部"},
		{1, "网站链接"},
		{2, "应用推广(ios)"},
		{3, "应用推广(android)"},
		{4, "小程序"},
		{5, "商品目录"},
		{6, "门店推广"},
		{7, "电商店铺"},
	}
}

func (e *FeedSubject) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = FeedSubject(v)
	return nil
}
