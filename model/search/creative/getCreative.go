package creative

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetCreativeRequest 查询推广创意 API Request
type GetCreativeRequest struct {
	// Ids 查询id集合;类型为单元ID时不超过1000个，类型为创意ID时不超过3000个，建议分批多次请求
	Ids []uint64 `json:"ids,omitempty"`
	// CreativeFields 查询推广创意字段
	// 取值范围：枚举值，列表如下
	// creativeId - 创意ID
	// adgroupId - 推广单元ID
	// title - 创意标题
	// pause - 暂停/启用创意
	// status - 创意状态
	// description1 - 创意描述第一行
	// description2 - 创意描述第二行
	// pcDestinationUrl - 计算机访问网址
	// pcDisplayUrl - 计算机显示网址
	// mobileDestinationUrl - 移动访问网址
	// mobileDisplayUrl - 移动显示网址
	// tabs - 标签
	// miniProgramUrl - 小程序访问网址
	// offlineReasons - 推广下线原因
	// deeplink - 应用调起网址
	// ulink - ios应用调起网址（该字段属于小流量功能，需要使用请向营销顾问申请）
	// createTime-添加时间
	// pcFinalUrl - 计算机最终访问网址
	// pcTrackParam - 计算机监控后缀
	// pcTrackTemplate - 计算机第三方追踪模板
	// mobileFinalUrl - 移动最终访问网址
	// mobileTrackParam - 移动监控后缀
	// mobileTrackTemplate - 移动第三方追踪模板
	CreativeFields []string `json:"creativeFields,omitempty"`
	// IdType 查询id类型;5 - 单元ID;7 - 创意ID
	IdType int `json:"idType,omitempty"`
	// GetTemp 是否获取创意影子;0 - 只查询创意本身;1 - 只查询创意影子;影子说明：用户先向系统提交了创意A，并且A已审核通过，之后再对A进行影响审核状态的修改（例如修改创意文案/url），修改后的创意为A’（A’即为影子，仅对审核通过的物料进行修改才会产生影子），在A’通过审核生效之前，线上的生效创意仍然为A。;此时：getTemp为0查询到的是A getTemp为1查询到的是A’
	GetTemp int `json:"getTemp,omitempty"`
}

func (r GetCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CreativeService/getCreative")
}

// GetCreativeResponse 获取推广创意 API Response
type GetCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
