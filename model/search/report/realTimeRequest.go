package report

// RealTimeRequest 实时报告请求-包含多类报告
type RealTimeRequest struct {
	// PerformanceData 必填 取值范围：cost（花费）、cpc（平均点击价格）、impression（展现）、click（点击）、ctr（点击率）、cpm（千次展现成本）、position（上方位平均排名）、conversion（网页转化）、商桥转化（bridgeConversion） 其中（impression,click必填） .另由于转化字段较多，取值可参考转化字段表。针对不同的物料层级，其合法的取值范围不同，（按照第一个请求的字段排序）请参见报告规则
	PerformanceData []string `json:"performanceData,omitempty"`
	// Order 是否为降序排列; 选填，默认null，按照时间排序： true：降序 false：升序 app下载报告/推广电话报告，不支持排序
	Order *bool `json:"order,omitempty"`
	// StartDate 统计开始时间，格式参考： 2010-08-01
	StartDate string `json:"startDate,omitempty"`
	// EndDate 统计结束时间，格式参考： 2010-08-17
	EndDate string `json:"endDate,omitempty"`
	// LevelOfDetails 指定返回的数据层级; 选填，默认为账户 2：账户粒度 3：计划粒度 5：单元粒度 7：创意粒度 11：关键词粒度 12：关键词+创意粒度 200: 高级创意报告粒度
	LevelOfDetails int `json:"levelOfDetails,omitempty"`
	// Attributes 针对特定的数据层级设置特定的统计范围; 选填； 为NULL表示统计全部地域。 key:provid； value:地域代码数组 说明：app下载报告/推广电话不支持attributes
	Attributes []Attribute `json:"attributes,omitempty"`
	// ReportType 实时数据类型; 必填； 2：账户 10：计划 11：单元 14：关键词 12：创意 3：地域 5：二级地域报告 40：app下载报告 928: 高级创意报告（目前仅支持未删除图片的数据）
	ReportType int `json:"reportType,omitempty"`
	// StatIds 统计范围下的id集合。根据StatRange的不同类型填写不同id	;选填，默认NULL，表示统计范围为全账户 最多500个 staRange为3时填写计划id; staRange为5时填写单元id; staRange为7时填写创意id; staRange为11时填写关键词keywordid
	StatIds []int64 `json:"statIds,omitempty"`
	// StatRange 统计范围; 选填，默认值为2； 2：账户范围 3：计划范围 5：单元范围 7：创意范围 11：关键词范围 注意：统计范围不能细于当前的统计粒度，例如统计粒度为计划，则统计范围不能细到单元
	StatRange int `json:"statRange,omitempty"`
	// UnitOfTime 统计时间单位;选填，默认值为5 取值范围： 5：分日 4：分周 3：分月 1：分年 7：分小时 8：请求时间段汇总(endDate-StartDate)
	UnitOfTime int `json:"unitOfTime,omitempty"`
	// Number 返回数据条数;选填 目前实时报告中账户、计划、单元、关键词、创意报告最大支持10000，其他类型实时报告只支持5000。 默认值1000 注意：超过限制或者小于等于0则报错 说明：app下载报告/推广电话报告、当物料量较大时，建议按计划或单元分批获取，一条对应三条返回值
	Number int `json:"number,omitempty"`
	// PageIndex 分页; 选填）请求页码 说明：默认不分页 当不填写页码或页码参数为<=0时，代表未分页； 请求页码设置大于0时，代表分页请求（同时需要配合number字段）
	PageIndex int `json:"pageIndex,omitempty"`
	// Device 搜索推广设备; 选填，默认值为0 取值范围： 0：全部搜索推广设备 1：仅计算机 2：仅移动
	Device int `json:"device,omitempty"`
}
