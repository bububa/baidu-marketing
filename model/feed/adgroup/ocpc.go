package adgroup

// Ocpc oCPC设置对象
type Ocpc struct {
	// AppTransId 转化ID; 营销目标为应用推广Android的单元，仅能绑定应用操作系统为Android的转化追踪；营销目标为应用推广iOS的单元，仅能绑定应用操作系统为iOS的转化追踪。
	AppTransId int64 `json:"appTransId,omitempty"`
	// TransFrom 接入方式; 1 - 应用API 2 - 基木鱼营销页; 4 - API激活; 5 - 网页JS布码; 7 - 线索API; 8 - 咨询工具授权; 13 - 应用SDK
	TransFrom int `json:"transFrom,omitempty"`
	// OcpcBid 目标转化出价; 取值范围：[0.3，999.99] 单位为元
	OcpcBid float64 `json:"ocpcBid,omitempty"`
	// LpUrl 推广URL; 必须以“http”或“https”开头
	LpUrl string `json:"lpUrl,omitempty"`
	// TransType 目标转化
	TransType int `json:"transType,omitempty"`
	// OcpcLevel oCPC投放阶段; 1 - 第一阶段; 2 - 第二阶段; 此字段仅用于查询，新增/修改推广单元时参数中传此字段无效
	OcpcLevel int `json:"ocpcLevel,omitempty"`
	// IsSkipStageOne 跳过数据积累; 仅当payMode=1时有效; 当取值为true时要求推广单元对象中bid=0，更新推广单元时不允许修改该字段
	IsSkipStageOne bool `json:"isSkipStageOne,omitempty"`
	// PayMode 付费模式
	PayMode int `json:"payMode,omitempty"`
	// OptimizeDeepTrans 优化深度转化;
	OptimizeDeepTrans *bool `json:"optimizeDeepTrans,omitempty"`
	// DeepOcpcBid 深度转化出价
	DeepOcpcBid float64 `json:"deepOcpcBid,omitempty"`
	// DeepTransType 深度转化类型
	DeepTransType int `json:"deepTransType,omitempty"`
	// UseRoi 使用ROI优化
	UseRoi *bool `json:"useRoi,omitempty"`
	// RoiRatio ROI转化率
	RoiRatio float64 `json:"roiRatio,omitempty"`
}
