package creative

import "github.com/bububa/baidu-marketing/model"

// Creative 推广创意
type Creative struct {
	// CampaignId 计划ID
	CampaignId int64 `json:"campaignId,omitempty"`
	// CreativeId 创意ID
	CreativeId int64 `json:"creativeId,omitempty"`
	// AdgroupId 推广单元ID
	AdgroupId int64 `json:"adgroupId,omitempty"`
	// Title 创意标题;长度限制：[9, 50];长度限制为字节数限制，1个中文按2个字节计算，英文、数字按1个字节计算，通配符不计入
	Title string `json:"title,omitempty"`
	// Description1 创意描述第一行;长度限制：[9, 80];长度限制为字节数限制，1个中文按2个字节计算，英文、数字按1个字节计算，通配符不计入
	Description1 string `json:"description1,omitempty"`
	// Description2 创意描述第二行;长度限制：[0, 80];长度限制为字节数限制，1个中文按2个字节计算，英文、数字按1个字节计算，通配符不计入
	Description2 string `json:"description2,omitempty"`
	// Pause 暂停/启用创意
	// 默认值：false
	// 取值范围：枚举值，列表如下
	// true - 暂停
	// false - 启用
	Pause *bool `json:"pause,omitempty"`
	// Status 创意状态
	// 取值范围：枚举值，列表如下
	// 51 - 有效
	// 52 - 暂停推广
	// 53 - 审核不通过
	// 54 - 待激活
	// 55 - 审核中
	// 56 - 部分无效
	// 57 - 有效-移动URL审核中
	// 59 - 未审核
	// 备注：部分无效：移动物料审核未通过，计算机物料审核通过
	Status int `json:"status,omitempty"`
	// MobileDestinationUrl 移动访问网址;长度限制：[1, 1024];网址域名需与账户注册域名相同，其他情况请参考业务限制
	MobileDestinationUrl *string `json:"mobileDestinationUrl,omitempty"`
	// MobileDisplayUrl 移动显示网址;长度限制：[1, 36];取值限制为账户注册域名本身，可通过"查询账户"接口获取，对应字段为regDomain
	MobileDisplayUrl *string `json:"mobileDisplayUrl,omitempty"`
	// PcDestinationUrl 计算机访问网址;长度限制：[1, 1024];网址域名需与账户注册域名相同，其他情况请参考业务限制
	PcDestinationUrl *string `json:"pcDestinationUrl,omitempty"`
	// PcDisplayUrl 计算机显示网址;长度限制：[1, 36];取值限制为账户注册域名本身，可通过"查询账户"接口获取，对应字段为regDomain
	PcDisplayUrl *string `json:"pcDisplayUrl,omitempty"`
	// OfflineReasons 推广下线原因;当有多个推广下线原因时，数组会有多个元素，每个代表一种原因
	OfflineReasons []model.OfflineReason `json:"offlineReasons,omitempty"`
	// Tabs 标签;说明: 同一创意支持添加多个标签，最多可添加30个，每个标签ID的取值范围为0-30，0表示无标签。;举例：创意同时标记了1、2、3号标签，则该字段取值为：[1,2,3]
	Tabs []int `json:"tabs,omitempty"`
	// Deeplink 应用调起网址;长度限制：[0, 1024];仅部分客户可设置，如需开通名单请咨询销售或客服同学
	Deeplink *string `json:"deeplink,omitempty"`
	// MiniProgramUrl 小程序访问网址;长度限制：[0, 1024]
	MiniProgramUrl *string `json:"miniProgramUrl,omitempty"`
	// PcFinalUrl 计算机最终访问网址
	PcFinalUrl *string `json:"pc_final_url,omitempty"`
	// PcTrackParam 计算机监控后缀
	PcTrackParam *string `json:"pc_track_param,omitempty"`
	// PcTrackTemplate 计算机第三方追踪模板
	PcTrackTemplate *string `json:"pc_track_template,omitempty"`
	// MobileFinalUrl 手机最终访问网址
	MobileFinalUrl *string `json:"mobile_final_url,omitempty"`
	// MobileTrackParam 手机监控后缀
	MobileTrackParam *string `json:"mobile_track_param,omitempty"`
	// MobileTrackTemplate 手机第三方追踪模板
	MobileTrackTemplate *string `json:"mobile_track_template,omitempty"`
}
