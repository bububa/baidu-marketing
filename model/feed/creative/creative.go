package creative

// Pause 可选值
// true - 暂停
// false - 启用

// Status 可选值
// 51 - 有效
// 52 - 暂停推广
// 53 - 审核不通过
// 54 - 待激活
// 55 - 审核中
// 56 - 部分无效
// 57 - 有效-移动URL审核中
// 备注：部分无效：移动物料审核未通过，计算机物料审核通过

type Creative struct {
	CreativeFeedId   int         `json:"creativeFeedId,omitempty"`
	AdgroupFeedId    int         `json:"adgroupFeedId,omitempty"`
	Materialstyle    []int       `json:"materialstyle,omitempty"`
	CreativeFeedName string      `json:"creativeFeedName,omitempty"`
	Pause            bool        `json:"pause,omitempty"`
	Status           []int       `json:"status,omitempty"`
	Material         string      `json:"material,omitempty"`
	Refusereason     string      `json:"refusereason,omitempty"`
	Playnum          []int       `json:"playnum,omitempty"`
	IdeaType         []int       `json:"ideaType,omitempty"`
	Addtime          interface{} `json:"addtime,omitempty"`
	Approvemsgnew    string      `json:"approvemsgnew,omitempty"`
	AuditTimeModel   struct {
		IdeaId           int         `json:"ideaId,omitempty"`
		IdeaState        []int       `json:"ideaState,omitempty"`
		StartTime        interface{} `json:"startTime,omitempty"`
		EstimationTime   interface{} `json:"estimationTime,omitempty"`
		EstimationMinute []int       `json:"estimationMinute,omitempty"`
		CompleteRadio    string      `json:"completeRadio,omitempty"`
	} `json:"auditTimeModel,omitempty"`
}
