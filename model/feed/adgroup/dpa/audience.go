package dpa

// Audience 商品定向设置
type Audience struct {
	// Premium 商品通投定向设置，必填
	Premium string `json:"premium,omitempty"`
	// 商品意图词定向设置，必填。
	PaKeywords string `json:"paKeywords,omitempty"`
	// Interets 商品长期兴趣定向设置，必填,定义同原生。
	Interets string `json:"interets,omitempty"`
	// NewInterets 商品新兴趣定向设置,名单内用户可用,定义同原生。
	NewInterets string `json:"newinterets,omitempty"`
	// PaCrowd 商品人群定向设置。
	PaCrowd string `json:"paCrowd,omitempty"`
	// Age 年龄。默认为全部，定义同原生。
	Age string `json:"age,omitempty"`
	// Sex 性别。默认为全部，定义同原生。
	Sex string `json:"sex,omitempty"`
	// Education 学历。默认为全部，定义同原生。
	Education string `json:"education,omitempty"`
	// Device 平台。默认为全部，定义同原生。
	Device string `json:"device,omitempty"`
	// Net 网络。默认为全部，定义同原生。
	Net string `json:"net,omitempty"`
	// Region 地域（省市）。默认为不限，定义同原生
	Region string `json:"region,omitempty"`
	// AutoRetion 智能地域。默认为关闭
	AutoRegion string `json:"autoRegion,omitempty"`
	// MediaCategoriesBindType  百青藤媒体分类定向方式，与mediaCategories字段配合使用。默认为不限，定义同原生
	MediaCategoriesBindType string `json:"mediaCategoriesBindType,omitempty"`
	// MediaCategories 百青藤媒体分类。默认为不限，定义同原生
	MediaCategories string `json:"mediaCategories"`
	// MediaIds 百青藤媒体ID。默认为不限，定义同原生
	MediaIds string `json:"mediaids,omitempty"`
	// MediaPackage 百青藤媒体包。默认为不限，不限用“”表示。目前仅支持设置一个媒体包,定义同原生
	MediaPackage string `json:"mediaPackage,omitempty"`
	// AutoExtend 自动扩量，原isOpenOcpcLab字段
	AutoExtend string `json:"autoExtend,omitempty"`
	// CustomMediaPackage 自定义媒体包，默认为不限，定义同原生
	CustomMediaPackage string `json:"customMediaPackage,omitempty"`
	// DeepLink 媒体包支持调起，仅流量类型为百青藤可用，默认false。定义同原生deeplinkOnly字段
	DeepLink string `json:"deepLink,omitempty"`
}

// PremiumAudience 商品通投定向设置
type PreminuAudience struct {
	Common bool `json:"common"`
}

// PaKeywordsAudience 商品意图词定向
type PaKeywordsAudience struct {
	// ShowWords 意图词数组，最多不能超过1000个，选择"意图词定向"时必填字段。每个意图词均为Object类型，参考下文商品意图词对象。
	ShowWords []PaKeyword `json:"showWords,omitempty"`
	// KeywordsExtend 用户行为，默认“1”。1: 全网行为; 2: 历史搜索; 当意图词数组中，包含行业词包和自定义词包类型的意图词时，用户行为只能设置为“全网行为”。
	KeywordsExtend int `json:"keywordsExtend,omitempty"`
	// KtFilter 无效意图词过滤，默认“False”。当意图词数组中，包含普通意图词时，过滤开关才生效。
	KtFilter bool `json:"ktFilter"`
	// AutoMatch 属否开启商品优选，默认“智能匹配”，其中“True”表示“智能匹配”。
	AutoMatch bool `json:"autoMatch"`
}

type PaKeyword struct {
	// Showword 意图词
	ShowWord string `json:"showword,omitempty"`
	// Type 意图词类型，默认"1"。1:普通意图词，最多不能超过900个; 2: 模版意图词，最多不能超过50个; 3: 行业词包，最多不能超过50个; 4: 自定义词包，最多不能超过1个
	Type int `json:"type,omitempty"`
}

// PaCrowd 商品人群定向
type PaCrowd struct {
	// Basic 基础人群，最多不超过10个。取值示例：[]或[121, 123, 145, …]
	Basic []int `json:"basic,omitempty"`
	// Common 交集人群，最多不超过10个。取值示例：[]或[121, 123, 145, …]
	Common []int `json:"common,omitempty"`
	// Exclude 排除人群，最多不超过10个。取值示例：[]或[121, 123, 145, …]
	Exclude []int `json:"exclude,omitempty"`
	// Redirect 是否开启商品重定向能力，默认“否”。只有当商品推广单元的商品基础定向为商品通投时，才会生效。
	Redirect bool `json:"redirect,omitempty"`
	// RtValidTime 商品重定向设置:时效性，1～90天，默认"90"。
	RtValidTime int `json:"rtValidTime,omitempty"`
	// RtWiseRecommend 商品重定向设置:智能推荐策略，默认“1”，请设置与商品重定向人群包相匹配的推荐策略
	RtWiseRecommend []int `json:"rtWiseRecommend,omitempty"`
	// RtBehaviors 商品重定向设置:行为类型，请设置与商品重定向人群包相匹配的推荐策略，默认"不限"，最多选择一种行为类型。
	RtBehaviors []int `json:"rtBehaviors,omitempty"`
}
