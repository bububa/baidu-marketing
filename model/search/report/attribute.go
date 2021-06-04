package report

type Attribute struct {
	Key   string `json:"key,omitempty"`   // 针对特定的报告类型设置特定的统计范围; 目前仅对地域报告和搜索词报告使用，合法值：provid
	Value []int  `json:"value,omitempty"` // 针对特定的报告类型设置特定的统计范围; 目前仅对地域报告和搜索词报告使用，key取provid时，value值为地域代码数组，为空则表示全部地域
}
