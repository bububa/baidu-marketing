package adgroup

// Adgroup 推广单元
type Adgroup struct {
	// AdgroupId 推广单元ID
	AdgroupId uint64 `json:"adgroupFeedId,omitempty"`
	// AdgroupName 推广单元名称
	AdgroupName string `json:"adgroupFeedName,omitempty"`
	// CampaignId 推广计划ID
	CampaignId uint64 `json:"campaignFeedId,omitempty"`
	// Pause 暂停/启用推广单元
	Pause *bool `json:"pause,omitempty"`
	// Status 推广单元状态
	// 取值范围：枚举值，列表如下
	// 0 - 有效
	// 1 - 暂停推广
	// 2 - 推广计划暂停推广
	// 3 - 直播结束后暂停
	// 4 - 转化追踪异常暂停
	// 5 - 基木鱼落地页暂停
	Status int `json:"status,omitempty"`
	// Audience 此字段为对象类型，对象内每个字段表示一种定向设置，所有字段的值均为String类型。每个单元可以设置0至多个定向设置。全部定向字段列表和详细说明请参考原生定向设置列表。
	Audience map[string]string `json:"audience,omitempty"`
	// Bid 出价，根据优化目标不同，具体含义如下：
	// CPC：单次点击出价
	// CPM：千次展现出价
	// oCPC：第一阶段单次点击出价
	// 出价为元，可精确到分。取值范围如下
	// 当优化目标选择CPC/oCPC，投放流量选择优选流量或者自定义流量且包含百青藤以外的流量时：[0.3,999.99]
	// 当优化目标选择CPC/oCPC，投放流量选择自定义流量且仅选择百青藤流量时：[0.2,999.99]
	// 当优化目标选择CPM时：[2.0,9999.99]
	Bid float64 `json:"bid,omitempty"`
	// FTypes 投放范围
	// 取值范围：枚举值，列表如下
	// 1 - 自定义类-百度信息流
	// 2 - 自定义类-贴吧
	// 4 - 百青藤（小程序营销目标-微信小游戏不支持）
	// 8 - 自定义类-好看视频（好看视频流量目前在实验阶段，仅限已开通该流量的账户使用）
	// 64 - 自定义类-百度小说
	// 空数组（[]）表示默认。默认、自定义、百青藤不可以同时选择
	FTypes []int `json:"ftypes,omitempty"`
	// BidType 优化目标和付费模式; 1 - 点击（CPC）(应用推广营销目标下不支持，不影响存量物料投放); 2 - 曝光（CPM）(应用推广营销目标下不支持，不影响存量物料投放); 3 - 转化（oCPC/oCPM）; 5 - eCPC
	BidType int `json:"bidtype,omitempty"`
	// Ocpc oCPC信息; 本字段只有当bidtype=3有效。请参考oCPC设置对象;
	Ocpc *Ocpc `json:"ocpc,omitempty"`
	// AtpFeedId 定向包ID; 定向包的优先级高于audience字段设置的定向信息。已绑定定向包单元解除绑定需传0
	AtpFeedId *int64 `json:"atpFeedId,omitempty"`
	// DeliveryType 投放场景
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 不限
	// 1 - 开屏
	// 2 - 激励
	// 4 - 原生
	// 开屏/激励/原生选项仅在投放范围选择百青藤时有效
	DeliveryType []int `json:"deliveryType,omitempty"`
	// UneffecientAdgroup 低效单元标识
	// 取值范围：枚举值，列表如下
	// 1 - 低效单元
	// 0 - 非低效单元
	UneffecientAdgroup int `json:"uneffecientAdgroup,omitempty"`
	// ProductSetId 商品组id；仅商品目录营销目标单元可用，新增时与unitProducts至少填写一个，与unitProducts同时使用时unitProducts优先生效。参照商品组管理。
	ProductSetId int64 `json:"productSetId,omitempty"`
	// UnitProducts 单元商品筛选设置；仅商品目录营销目标单元可用，新增时与productSetId至少填写一个，与productSetId同时使用时unitProducts优先生效
	UnitProducts []UnitProducts `json:"unitProducts,omitempty"`
	// FtypeSelection 流量来源
	// 取值范围：枚举值，列表如下
	// 1 - 单元单独设置流量
	// 2 - 使用计划流量设置
	// 使用计划流量设置时会忽略传参中的ftypes，优先使用计划流量
	// 单元单独设置流量时单元流量需要为计划流量的子集
	// 出价上移名单使用字段，名单外使用无效。
	FtypeSelection int `json:"ftypeSelection,omitempty"`
	// BidSource 出价来源
	// 取值范围：枚举值，列表如下
	// 1 - 单元单独设置出价
	// 2 - 使用计划出价
	// 使用计划出价时，不需要传bid、bidtype以及oCPC中的出价相关字段，优先使用计划出价设置
	// 出价上移名单使用字段，名单外使用无效。
	BidSource int `json:"bidSource,omitempty"`
	// UrlType 落地页类型
	// 取值范围：枚举值，列表如下
	// 1 - 普通落地页
	// 2 - 百度小程序
	// 3 - 直播间
	// 使用计划出价后，落地页类型可以直接使用urlType字段，优先级高于oCPC字段
	UrlType int `json:"urlType,omitempty"`
	// MiniProgram 小程序信息
	// 该字段是json字符串类型
	// 使用计划出价后，落地页类型为百度小程序时miniProgramType、appKey、pagePath可以直接使用miniProgram，优先级高于oCPC字段
	MiniProgram string `json:"miniProgram,omitempty"`
	// BroadCastInfo 直播间信息
	// 该字段是json字符串类型
	// 使用计划出价后，落地页类型为直播间时broadCastMode、anchorId、roomId可以直接使用broadCastInfo，优先级高于oCPC字段
	BroadCastInfo string `json:"broadCastInfo,omitempty"`
	// Url 落地页
	// 使用计划出价后，落地页可以直接使用url字段，优先级高于oCPC字段；
	// 营销目标为小程序，子类型为微信小游戏时，仅支持基木鱼落地页
	Url string `json:"url,omitempty"`
	// UnitOcpxStatus 	单元学习状态
	// 取值范围：枚举值，列表如下
	// 1 - 正在学习
	// 2 - 学习成功
	// 3 - 学习失败
	UnitOcpxStatus int `json:"unitOcpxStatus,omitempty"`
	// CreateAtp 是否创建定向包
	// 默认为false
	// 取值范围：枚举值，列表如下
	// true - 创建定向包
	// false - 不创建定向包
	CreateAtp *bool `json:"createAtp,omitempty"`
	// AtpName 定向包名称
	AtpName string `json:"atpName,omitempty"`
	// AtpDesc 定向包描述
	AtpDesc string `json:"atpDesc,omitempty"`
}

// UnitProducts 单元商品筛选设置
type UnitProducts struct {
	// CatalogId 商品目录id；商品目录管理请登录百度商品中心
	CatalogId uint64 `json:"catalogId,omitempty"`
	// RuleProducts 商品组规则
	RuleProducts []ProductSetRule `json:"ruleProducts,omitempty"`
}

// ProductSetRule 商品组规则
type ProductSetRule struct {
	// Field 字段名；参照查询商品库字段返回的字段value
	Field string `json:"field,omitempty"`
	// Operator 操作符。
	// 取值范围：枚举值，列表如下
	// EQUAL - 等于
	// NOT_EQUAL - 不等于
	// CONTAIN - 包含
	// NOT_CONTAIN - 不包含
	// GREATER - 大于
	// LESS - 小于
	// GREATER_EQUAL - 大于等于
	// LESS_EQUAL - 小于等于
	Operator string `json:"operator,omitempty"`
	// Values 操作值。 操作字段值，多个操作值使用'|'分隔，多个之间取并集，最多15个值
	Values []string `json:"values,omitempty"`
}
