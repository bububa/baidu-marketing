package report

type AdgroupReport struct {
	*AdStats
	Date               string `json:"date"`                         // 日期
	UserName           string `json:"userName,omitempty"`           // 账户
	CampaignNameStatus string `json:"campaignNameStatus,omitempty"` // 推广计划
	AdGroupNameStatus  string `json:"adGroupNameStatus,omitempty"`  // 推广单元
	FeedFlowTypeEnum   int    `json:"feedFlowTypeEnum,omitempty"`   // 流量类型
	AdGroupStatus      string `json:"adGroupStatus,omitempty"`      // 推广单元状态
	AdGroupName        string `json:"adGroupName,omitempty"`        // 推广单元
	CampaignID         int64  `json:"campaignId,omitempty"`         // 推广计划ID
	AdGroupID          int64  `json:"adGroupId,omitempty"`          // 推广单元ID
	CampaignStatus     string `json:"campaignStatus,omitempty"`     // 推广计划状态
	CampaignName       string `json:"campaignName,omitempty"`       // 推广计划
	Device             int    `json:"device,omitempty"`             // 推广设备
	FeedSubjectEnum    int    `json:"feedSubjectEnum,omitempty"`    // FeedSubject
	BsType             int    `json:"bsType,omitempty"`             // 报告类型

	OCPCTransType                  string  `json:"ocpcTransType,omitempty"`                  // OCPCTransType 表示 OCPC 广告的目标转化类型。 枚举值参考接口文档说明。
	OCPCTargetTrans                int64   `json:"ocpcTargetTrans,omitempty"`                // OCPCTargetTrans 表示目标转化量。
	OCPCTargetTransCPC             float64 `json:"ocpcTargetTransCPC,omitempty"`             // OCPCTargetTransCPC 表示目标转化成本。
	OCPCTargetTransRatio           float64 `json:"ocpcTargetTransRatio,omitempty"`           // OCPCTargetTransRatio 表示转化率。
	OCPCDeepTargetTrans            int64   `json:"ocpcDeepTargetTrans,omitempty"`            // OCPCDeepTargetTrans 表示深度转化量。
	OCPCDeepTargetTransCPC         float64 `json:"ocpcDeepTargetTransCPC,omitempty"`         // OCPCDeepTargetTransCPC 表示深度转化成本。
	OCPCDeepTargetTransRatio       float64 `json:"ocpcDeepTargetTransRatio,omitempty"`       // OCPCDeepTargetTransRatio 表示深度转化率。
	ConvertTimeOCPCTargetTrans     int64   `json:"convertTimeOCPCTargetTrans,omitempty"`     // ConvertTimeOCPCTargetTrans 表示目标转化量的转化时间。
	ConvertTimeOCPCDeepTargetTrans int64   `json:"convertTimeOCPCDeepTargetTrans,omitempty"` // ConvertTimeOCPCDeepTargetTrans 表示深度转化量的转化时间。
	DeepConvertType                string  `json:"deepConvertType,omitempty"`                // DeepConvertType 表示 OCPC 广告的深度转化类型。 枚举值参考接口文档说明。
}
