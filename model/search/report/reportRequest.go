package report

// 异步报告请求
type ReportRequest struct {
	PerformanceData []string    `json:"performanceData,omitempty"` // 必填 取值范围：cost（花费）、cpc（平均点击价格）、impression（展现）、click（点击）、ctr（点击率）、cpm（千次展现成本）、position（上方位平均排名）、conversion（网页转化）、商桥转化（bridgeConversion） 其中（impression,click必填） .另由于转化字段较多，取值可参考转化字段表。针对不同的物料层级，其合法的取值范围不同，（按照第一个请求的字段排序）请参见报告规则
	IdOnly          bool        `json:"idOnly,omitempty"`          // 是否只需要id; 选填；默认为false 取值范围： true：只获取id false：既获取id也获取字面
	StartDate       string      `json:"startDate,omitempty"`       // 统计开始时间，格式参考： 2010-08-01
	EndDate         string      `json:"endDate,omitempty"`         // 统计结束时间，格式参考： 2010-08-17
	LevelOfDetails  int         `json:"levelOfDetails,omitempty"`  // 指定返回的数据层级; 选填，默认为账户 2：账户粒度 3：计划粒度 5：单元粒度 7：创意粒度 11：关键词粒度 12：关键词+创意粒度 200: 高级创意报告粒度
	Attributes      []Attribute `json:"attributes,omitempty"`      // 针对特定的数据层级设置特定的统计范围; 选填； 为NULL表示统计全部地域。 key:provid； value:地域代码数组 说明：app下载报告/推广电话不支持attributes
	Format          int         `json:"format,omitempty"`          // 报告文件格式; 选填，默认值为2； 2：csv格式
	ReportType      int         `json:"reportType,omitempty"`      // 实时数据类型; 必填； 2：账户 10：计划 11：单元 14：关键词 12：创意 3：地域 5：二级地域报告 40：app下载报告 928: 高级创意报告（目前仅支持未删除图片的数据）
	StatIds         []int64     `json:"statIds,omitempty"`         // 统计范围下的id集合。根据StatRange的不同类型填写不同id	;选填，默认NULL，表示统计范围为全账户 最多500个 staRange为3时填写计划id; staRange为5时填写单元id; staRange为7时填写创意id; staRange为11时填写关键词keywordid
	StatRange       int         `json:"statRange,omitempty"`       // 统计范围; 选填，默认值为2； 2：账户范围 3：计划范围 5：单元范围 7：创意范围 11：关键词范围 注意：统计范围不能细于当前的统计粒度，例如统计粒度为计划，则统计范围不能细到单元
	UnitOfTime      int         `json:"unitOfTime,omitempty"`      // 统计时间单位;选填，默认值为5 取值范围： 5：分日 4：分周 3：分月 1：分年 7：分小时 8：请求时间段汇总(endDate-StartDate)
	Device          int         `json:"device,omitempty"`          // 搜索推广设备; 选填，默认值为0 取值范围： 0：全部搜索推广设备 1：仅计算机 2：仅移动
}
