package report

// ReportRequest 实时报告请求-包含多类报告
// LevelOfDetails 类型
// 2：账户粒度
// 3：计划粒度
// 5：单元粒度
// 7：创意粒度
//
// ReportType 类型
// 700：账户报告
// 701：计划报告
// 702：单元报告
// 703：创意报告
// 706：地域报告
// 754：小程序报告
// 730：意图词报告
// 725：ocpx报告
// 757：媒体ID报告（媒体ID报告目前有权限限制，仅限开通媒体ID定向的账户使用）
// 705：程序化创意组合报告
// 711：性别年龄学历报告
// 717：平台网络报告
// 756：商圈报告
// 718：场所报告
// 747：APP分类报告
// 748：APP名称报告
// 751：人生阶段报告
// 745：文章类型报告
// 2371731：自动扩量报告
//
// StatRange 可选值
// 选填，默认值为2
// 2：账户范围：
// 3：计划范围
// 5：单元范围
// 7：创意范围
// 注意：统计范围不能细于当前的统计粒度，例如统计粒度为计划，则统计范围不能细到单元
//
// UnitOfTime 可选值
// 选填，默认值为5
// 取值范围：
// 5：分日
// 4：分周
// 3：分月
// 7：分时
// 8：请求时间段汇总(endDate-StartDate)
//
// Subject 可选值
// 选填，默认值为0
// 0：全部
// 1：网站链接
// 2：应用下载（IOS）
// 3：应用下载（ANDROID）
// 注意：账户报告不支持此字段
//
// BidType 可选值
// ocpx才有用
// ["3"]：ocpx
// ["5"]：ecpc
// ["3"，"5"]：ocpx＋ecpc
//
// PayMode 可选值
// ocpx才有用
// ["0","1"]：ocpc
// ["2"]：ocpm
// ["0","1","2"]：ocpx
//
// MaterialStyle 可选值
// 选填，默认值为0
// 0：全部
// 1：单图
// 2：三图
// 3：大图
// 4：视频
// 5：橱窗
// 6：开屏
// 7：横幅
// 注意：仅创意报告支持此字段
//
// bstyle 可选值
// 选填（不填返回全部）
// 0：全部
// 1：普通报告
// 3：闪投报告
// 7：原生RTA
type ReportRequest struct {
	// PerformanceData 必填 取值范围：cost（花费）、cpc（平均点击价格）、impression（展现）、click（点击）、ctr（点击率）、cpm（千次展现成本）、position（上方位平均排名）、conversion（网页转化）、商桥转化（bridgeConversion） 其中（impression,click必填） .另由于转化字段较多，取值可参考转化字段表。针对不同的物料层级，其合法的取值范围不同，（按照第一个请求的字段排序）请参见报告规则
	PerformanceData []string `json:"performanceData,omitempty"`
	// IdOnly 是否只需要id
	IdOnly *bool `json:"idOnly,omitempty"`
	// StartDate 统计开始时间，格式参考： 2010-08-01
	StartDate string `json:"startDate,omitempty"`
	// EndDate 统计结束时间，格式参考： 2010-08-17
	EndDate string `json:"endDate,omitempty"`
	// LevelOfDetails 指定返回的数据层级;
	LevelOfDetails int `json:"levelOfDetails,omitempty"`
	// Format 报告文件格式
	Format int `json:"format,omitempty"`
	// ReportType 实时数据类型;
	ReportType int `json:"reportType,omitempty"`
	// StatIds 统计范围下的id集合。根据StatRange的不同类型填写不同id;选填，默认NULL，表示统计范围为全账户;staRange为3时填写计划id;staRange为5时填写单元id;staRange为7时填写创意id
	StatIds []int64 `json:"statIds,omitempty"`
	// StatRange 统计范围;
	StatRange int `json:"statRange,omitempty"`
	// UnitOfTime 统计时间单位;
	UnitOfTime int `json:"unitOfTime,omitempty"`
	// Producttype 版位; 0-全部; 1-列表页; 2-详情页
	Producttype int `json:"producttype,omitempty"`
	// Subject 推广对象
	Subject int `json:"subject,omitempty"`
	// BidType 出价模式
	BidType []string `json:"bidType,omitempty"`
	// PayMode 付费模式
	PayMode []string `json:"payMode,omitempty"`
	// MaterialStyle 样式类型
	MaterialStyle int `json:"materialStyle,omitempty"`
	Bstype        int `json:"bstype,omitempty"` // Bstype bstype;
}
