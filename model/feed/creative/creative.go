package creative

// Creative 创意
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
	// CreativeFeedId 创意ID
	CreativeFeedId int `json:"creativeFeedId,omitempty"`
	// AdgroupFeedId 所属推广单元ID
	AdgroupFeedId int `json:"adgroupFeedId,omitempty"`
	// Materialstyle 创意样式ID
	Materialstyle int `json:"materialstyle,omitempty"`
	// CreativeFeedName 创意名称
	CreativeFeedName string `json:"creativeFeedName,omitempty"`
	// Pause 是否暂停推广
	Pause bool `json:"pause,omitempty"`
	// Status 创意状态
	Status int `json:"status,omitempty"`
	// Material 物料内容
	Material string `json:"material,omitempty"`
	// Refusereason 审核未通过的原因（审核拒绝理由）
	Refusereason string `json:"refusereason,omitempty"`
	// Playnum 视频创意的视频播放量
	Playnum int `json:"playnum,omitempty"`
	// IdeaType 创意类型
	IdeaType int `json:"ideaType,omitempty"`
	// Addtime 创意的创建时间，仅获取创意列表时传递有效
	Addtime string `json:"addtime,omitempty"`
	// Approvemsgnew JSON格式化后的审核未通过的原因（审核拒绝理由）
	Approvemsgnew string `json:"approvemsgnew,omitempty"`
	// AuditTimeModel 预估审核时间
	AuditTimeModel *AuditTimeModel `json:"auditTimeModel,omitempty"`
}

// AuditTimeModel 预估审核时间
type AuditTimeModel struct {
	// IdeaId 审核ID
	IdeaId int `json:"ideaId,omitempty"`
	// IdeaState 审核状态
	IdeaState []int `json:"ideaState,omitempty"`
	// StartTime 审核开始时间
	StartTime string `json:"startTime,omitempty"`
	// EstimationTime 预计审核结束时间
	EstimationTime string `json:"estimationTime,omitempty"`
	// EstimationMinute 预计审核时长
	EstimationMinute []int `json:"estimationMinute,omitempty"`
	// CompleteRatio 完成百分比
	CompleteRadio string `json:"completeRadio,omitempty"`
}
