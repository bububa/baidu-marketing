package model

type OfflineReason struct {
	MainReason   string `json:"mainReason,omitempty"`   // 推广下线主要原因ID，值为”3”时，代表审核不通过
	DetailReason string `json:"detailReason,omitempty"` // 推广下线具体原因，当mainReason为“3”时，本字段代表审核不通过的具体原因;结构说明：detailReason为json字符串，解析后为嵌套数组，内层数组第一个元素为具体拒绝理由;处理说明：字符串先解析成外层json数组，取外层数组第一个元素json数组作为内层数组，取内层数组的第一个元素字符串;示例："[["您提交的物料涉及不合规内容，可能涉及以下问题（物料文字不合规:前面有名词但不是主语），请修改提交内容",""]]";
}
