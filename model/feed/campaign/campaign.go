package campaign

import "github.com/bububa/baidu-marketing/model"

// Campaign 计划对象
type Campaign struct {
	// CampaignId 计划ID
	CampaignId uint64 `json:"campaignFeedId,omitempty"`
	// CampaignName 计划名称。长度限制最大100个字节，1个中文及中文符号按2个字节计算
	CampaignName string `json:"campaignFeedName,omitempty"`
	// Subject 推广对象
	// 取值范围：枚举值，列表如下
	// 1 - 网站链接
	// 2 - 应用下载（IOS）
	// 3 - 应用下载（Android）
	// 4 - 小程序（需要开通小流量名单，目前仅支持微信小游戏）
	// 7 - 电商店铺
	// 8 - 销售线索
	// 9- 应用调起
	Subject int `json:"subject,omitempty"`
	// AppInfo 推广app信息。subject=1时，该字段无效。对象定义参考下文推广app信息
	AppInfo *AppInfo `json:"appInfo,omitempty"`
	// Budget 推广计划预算。默认为0,表示不限预算。正常取值范围为[50 - 9999999.99]
	Budget float64 `json:"budget,omitempty"`
	// StartTime 推广开始日期。默认为null，表示长期投放。格式示例：'2016-12-15'不能早于当天的日期
	StartTime string `json:"starttime,omitempty"`
	// EndTime 推广结束日期。默认为null，表示长期投放。例如：'2016-12-18'不能早于开始日期
	EndTime string `json:"endtime,omitempty"`
	// Schedule 暂停时段设置，对象定义参考下文暂停时段设置
	Schedule []model.Schedule `json:"schedule,omitempty"`
	// Pause 是否暂停推广。默认为false。true：推广计划暂停 false：推广计划启用
	Pause *bool `json:"pause,omitempty"`
	// Status 推广计划状态。
	// 取值范围：枚举值，列表如下
	// 0 - 有效
	// 1 - 处于暂停时段
	// 2 - 暂停推广
	// 3 - 推广计划预算不足
	// 4 - 账户待激活
	// 11 - 账户预算不足
	// 20 - 账户余额为零
	// 23 - 被禁推
	// 24 - app已下线
	// 25 - 应用审核中
	// 26 - RTA计划暂停
	// 27 - 计划绑定的新游预约包预约过期
	// 28 - 项目暂停
	Status int `json:"status,omitempty"`
	// BsType 物料类型
	// 1：普通计划 3：闪投计划 7：原生RTA 注：不支持修改
	BsType int `json:"bstype,omitempty"`
	// CampaignType 信息流计划类型。
	// 1： 普通模式 4：放量模式
	CampaignType int `json:"campaignType,omitempty"`
	// AddTime 添加时间
	AddTime string `json:"addtime,omitempty"`
	// EshopType 交易所在平台。
	// 取值范围：枚举值，列表如下
	// 1 - 基木鱼
	// 3-1 - 淘宝店铺（含天猫）
	// 3-2 - 京东店铺
	// 3-3 - 拼多多店铺
	// 3-4 - 苏宁店铺
	// 4 - 百度健康商城
	// 仅推广对象为电商店铺时需传该字段
	// 交易平台为基木鱼（度小店）、百度健康商城需要使用电商广告投放接口进行新建投放
	EshopType string `json:"eshopType,omitempty"`
	// Shadow 应用推广营销目标下的影子计划
	Shadow *AppInfoShadow `json:"shadow,omitempty"`
	// BudgetOfflineTime 当天计划预算下线最近一次的时间
	BudgetOfflineTime string `json:"budgetOfflineTime,omitempty"`
	// RtaStatus RTA状态
	// 取值范围：枚举值，列表如下
	// 1 - 已开通
	// 0 - 未开通
	RtaStatus int `json:"rtaStatus,omitempty"`
	// FtTypes 流量类型
	// 取值范围：枚举值，列表如下
	// 1 - 自定义类-百度信息流
	// 2 - 自定义类-贴吧
	// 4 - 百青藤（小程序营销目标-微信小游戏不支持）
	// 8 - 自定义类-好看视频（好看视频流量目前在实验阶段，仅限已开通该流量的账户使用）
	// 64 - 自定义类-百度小说
	// 空数组（[]）表示默认。默认、自定义、百青藤不可以同时选择，出价上移名单使用字段，名单外使用无效。
	FtTypes []int `json:"ftTypes,omitempty"`
	// BidType 出价方式
	// 默认值：1
	// 取值范围：枚举值，列表如下
	// 1 - 点击（CPC）
	// 2 - 曝光（CPM）
	// 3 - 转化（oCPC/oCPM） 出价上移名单使用字段，名单外使用无效。
	BidType int `json:"bidType,omitempty"`
	// Bid 出价，根据优化目标不同，具体含义如下：
	// CPC：单次点击出价
	// CPM：千次展现出价
	// oCPC：第一阶段单次点击出价
	// 出价为元，可精确到分。取值范围如下
	// 当优化目标选择CPC/oCPC，投放流量选择优选流量或者自定义流量且包含百青藤以外的流量时：[0.3,999.99]
	// 当优化目标选择CPC/oCPC，投放流量选择自定义流量且仅选择百青藤流量时：[0.2,999.99]
	// 当优化目标选择CPM时：[2.0,9999.99]
	// 出价上移名单使用字段，名单外使用无效。
	Bid float64 `json:"bid,omitempty"`
	// Ocpc oCPC相关设置
	// 本字段只有当bidtype=3有效，出价上移名单使用字段，名单外使用无效。
	Ocpc *OcpcModel `json:"ocpc,omitempty"`
	// UneffecientCampaign 低效计划标识
	// 取值范围：枚举值，列表如下
	// 1 - 低效计划
	// 0 - 非低效计划
	UneffecientCampaign int `json:"uneffecientCampaign,omitempty"`
	// BmcUserId 商品中心用户ID
	// 请使用“商品中心账号列表” 接口获取用户ID
	BmcUserId uint64 `json:"bmcUserId,omitempty"`
	// CategoryId 	商品目录ID产品分组ID,根据catalogSource字段决定
	// 请使用“产品组/商品目录列表” 接口获取商品目录ID
	CategoryId uint64 `json:"categoryId,omitempty"`
	// CatalogSource 	商品目录来源
	// 取值范围：枚举值，列表如下
	// 1 - 商品中心(BMC)
	// 2 - 基木鱼内容中心
	CatalogSource int `json:"catalogSource,omitempty"`
	// ProductType 产品库类型
	// 取值范围：枚举值，列表如下
	// 1 - 小说库
	// 2 - 短剧库
	// 3 - 长视频库
	// 4 - 课程中心
	// 5 - 医美项目
	// 6 - 零售库
	// 7 - 资讯库
	// 10001 - 内容中心
	// 关联产品库时必填
	ProductType int `json:"productType,omitempty"`
	// ProjectFeedId 项目ID
	// 传0表示解绑
	ProjectFeedId *uint64 `json:"projectFeedId,omitempty"`
	// InheritAscriptionType 继承归属类型
	// 取值范围：枚举值，列表如下
	// 1 - 当前账户
	// 2 - 同客户中心
	// 3 - 当前账户内的计划
	// 4 - 同客户中心的计划
	InheritAscriptionType int `json:"inheritAscriptionType,omitempty"`
	// InheritUserIds 继承优质计划账户ID集合
	// 当inheritAscriptionType为1时，只能传当前账户ID
	// 当inheritAscriptionType为2时需要传同客户中心的账户ID，数量限制为100
	InheritUserIds []uint64 `json:"inheritUserIds,omitempty"`
	// InheritCampaignInfos 继承计划信息集合
	// 当inheritAscriptionType为3时，只能传当前账户下的计划，计划数量限制为100个
	// 当inheritAscriptionType为4时需要传同客户中心的账户ID，单个账户下的计划数量限制为100个
	InheritCampaignInfos []InheritCampaignInfo `json:"inheritCampaignInfos,omitempty"`
	// CampaignOcpxStatus 计划学习状态
	// 取值范围：枚举值，列表如下
	// 1 - 正在学习
	// 2 - 学习成功
	// 3 - 学习失败
	CampaignOcpxStatus int `json:"campaignOcpxStatus,omitempty"`
	// UseLiftBudget 是否开启一键起量
	// 取值范围：枚举值，列表如下
	// 0 - 不开启
	// 1 - 开启
	UseLiftBudget *int `json:"useLiftBudget,omitempty"`
	// LiftBudget 起量预算
	// 起量预算最小值设置需要为出价10倍或计划预算的20%，最大值不能超过计划预算
	LiftBudget float64 `json:"liftBudget,omitempty"`
	// LiftStatus 起量状态
	// 取值范围：枚举值，列表如下
	// 0 - 未开启起量预算
	// 1 - 起量中
	// 2 - 起量完成
	LiftStatus int `json:"liftStatus,omitempty"`
	// DeliveryType 投放场景
	// 取值范围：枚举值，列表如下
	// 0 - 不限
	// 1 - 开屏
	// 2 - 激励
	// 4 - 原生
	// 开屏/激励/原生仅在投放范围选择百青藤时有效
	DeliveryType []int `json:"deliveryType,omitempty"`
	// AppSubType 应用推广子类型
	// 取值范围：枚举值，列表如下
	// 0 - 应用下载
	// 1 - 新游预约
	// 2 - 应用调起
	AppSubType *int `json:"appSubType,omitempty"`
	// MiniProgramType 小程序子类型
	// 取值范围：枚举值，列表如下
	// 1 - 百度小程序
	// 3 - 微信小程序
	// 4 - 微信小游戏
	MiniProgramType *int `json:"miniProgramType,omitempty"`
	// BidMode 出价模式
	// 取值范围：枚举值，列表如下
	// 1 - 目标转化成本
	BidMode int `json:"bidMode,omitempty"`
	// ProductIds 产品ID
	// 计划加入项目后，在项目上关联的产品ID
	ProductIds string `json:"productIds,omitempty"`
}
