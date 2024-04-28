package enum

// FeedMaterialStyle 该字段仅支持用于Filter中筛选数据使用。
type FeedMaterialStyle Enum[int]

func (e FeedMaterialStyle) Values() []FeedMaterialStyle {
	return []FeedMaterialStyle{
		{0, "全部"},
		{1, "单图"},
		{10, "互动图"},
		{100, "程序化"},
		{16, "开屏视频"},
		{2, "三图"},
		{20, "全部视频"},
		{3, "大图"},
		{5, "橱窗"},
		{6, "开屏图片"},
		{7, "横幅"},
		{8, "横版视频"},
		{9, "竖版视频"},
	}
}

func (e *FeedMaterialStyle) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = FeedMaterialStyle(v)
	return nil
}
