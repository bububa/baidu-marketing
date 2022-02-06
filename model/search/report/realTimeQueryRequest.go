package report

// RealTimeQueryRequest 搜索词报告请求
// querystatus（用户添加状态）返回值与字面映射关系
// "0": "已添加",
// "1": "未添加",
// "2": "不可添加"
// mixWmatchEnum（匹配模式）返回值与字面映射关系
// "0": "智能匹配"
// "15": "智能匹配"
// "127": "分匹配出价"
// "16": "智能匹配核心词"
// "313": "智能匹配核心词"
// "17": "短语匹配"
// "18": "短语匹配"
// "31": "短语匹配"
// "48": "精确匹配"
// "63": "精确匹配"
// "9998":"-"
// "9999":"-"
type RealTimeQueryRequest struct {
	// PerformanceData 必填 取值范围：click（点击）、 impression（展现）、cost 花费）、ctr（点击率）、trans（网页转化）、bridgetrans（商桥转化）、querystatus（账户添加状态）、mixWmatchEnum（匹配模式）（impression,click必填） 针对不同的物料层级，其合法的取值范围不同，（按照第一个请求的字段排序）请参见规则描述。
	PerformanceData []string `json:"performanceData,omitempty"`
	// StartDate 统计开始时间，格式参考： 2010-08-01
	StartDate string `json:"startDate,omitempty"`
	// EndDate 统计结束时间，格式参考： 2010-08-17
	EndDate string `json:"endDate,omitempty"`
	// LevelOfDetails 指定返回的数据层级; 选填，默认为账户 2：账户粒度 3：计划粒度 5：单元粒度 7：创意粒度 11：关键词粒度 12：关键词+创意粒度 200: 高级创意报告粒度
	LevelOfDetails int `json:"levelOfDetails,omitempty"`
	// Attributes 针对特定的数据层级设置特定的统计范围; 选填； 为NULL表示统计全部地域。 key:provid； value:地域代码数组 说明：app下载报告/推广电话不支持attributes
	Attributes []Attribute `json:"attributes,omitempty"`
	// ReportType 实时数据类型; 必填； 6：搜索词
	ReportType int `json:"reportType,omitempty"`
	// StatIds 统计范围下的id集合。根据StatRange的不同类型填写不同id	;选填，默认NULL，表示统计范围为全账户 staRange为3时填写计划id;
	StatIds []int64 `json:"statIds,omitempty"`
	// StatRange 统计范围; 选填，默认值为2； 2：账户范围 3：计划范围 注意：统计范围不能细于当前的统计粒度，例如统计粒度为计划，则统计范围不能细到单元
	StatRange int `json:"statRange,omitempty"`
	// UnitOfTime 统计时间单位; 选填，默认值为5 取值范围： 5：分日 8：请求时间段汇总(endDate-StartDate)
	UnitOfTime int `json:"unitOfTime,omitempty"`
	// Number 返回数据条数;选填 目前实时报告中账户、计划、单元、关键词、创意报告最大支持10000，其他类型实时报告只支持5000。 默认值1000 注意：超过限制或者小于等于0则报错 说明：app下载报告/推广电话报告、当物料量较大时，建议按计划或单元分批获取，一条对应三条返回值
	Number int `json:"number,omitempty"`
	// Device 搜索推广设备; 选填，默认值为0 取值范围： 0：全部搜索推广设备 1：仅计算机 2：仅移动
	Device int `json:"device,omitempty"`
}
