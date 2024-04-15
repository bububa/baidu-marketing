package report

// Stat 统计
type Stat struct {
	Dimension
	Metric
}

// Dimension 维度
type Dimension struct {
	// Date 日期
	// 不同时间单位对应的日期格式：
	// DAY: 2021-03-31
	// HOUR: 2021-03-31 23:00
	// WEEK MONTH SUMMARY: 2021-03-01至2021-03-31，按照时间单位对应的周期汇总
	Date string `json:"date,omitempty"`
	// UserId 账户ID
	UserId uint64 `json:"userId,omitempty"`
	// UserName 账户
	UserName string `json:"userName,omitempty"`
	// BusinessPointId 推广业务ID
	BusinessPointId uint64 `json:"businessPointId,omitempty"`
	// BusinessPointName 推广业务
	BusinessPointName string `json:"businessPointName,omitempty"`
	// WBudget 预算
	WBudget float64 `json:"wBudget,omitempty"`
	// BudgetEffeciency 预算利用率
	BudgetEffeciency float64 `json:"budgetEffeciency,omitempty"`
	// CampaignNameStatus 推广计划
	// 报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"推广计划[已删除]"。
	// 未删除物料正常显示名称，没有额外标记。
	CampaignNameStatus string `json:"campaignNameStatus,omitempty"`
	// AdGroupNameStatus 推广单元
	// 报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"推广单元[已删除]"。
	// 未删除物料正常显示名称，没有额外标记。
	AdGroupNameStatus string `json:"adGroupNameStatus,omitempty"`
	// AdGroupStatus 推广单元状态
	// 0：未删除，1：已删除
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:0
	// 1:1
	AdgroupStatus int `json:"adgroupStatus,omitempty"`
	// AdGroupName 推广单元
	AdGroupName string `json:"adGroupName,omitempty"`
	// CampaignId 推广计划ID
	CampaignId uint64 `json:"campaignId,omitempty"`
	// AdGroupId 推广单元ID
	AdGroupId uint64 `json:"adGroupId,omitempty"`
	// CampaignStatus 推广计划状态
	// 0：未删除，1：已删除
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:0
	// 1:1
	CampaignStatus int `json:"campaignStatus,omitempty"`
	// CampaignName 推广计划
	CampaignName string `json:"campaignName,omitempty"`
	// IdeaId 创意ID
	IdeaId uint64 `json:"ideaId,omitempty"`
	// IdeaInfo 创意
	// 必填，组合列，包含如下属性：
	// ideaId：创意ID
	// ideaTitle：创意标题
	// ideaTitleStatus：创意标题
	// ideaShowURL：显示URL
	// ideaMShowURL：显示URL(移动)
	// ideaDesc：创意描述1
	// ideaDesc2：创意描述2
	// 组合列在返回数据时会放到一个内嵌的Map结构中，比如：
	// {"ideaInfo":{"ideaTitle":"xxx","ideaDesc":"xxx","ideaTitleStatus":"xxx","ideaMShowURL":"xxx","ideaDesc2":"xxx","ideaShowURL":"xxx","ideaId":"xxx"}}
	IdeaInfo *IdeaInfo `json:"ideaInfo,omitempty"`
	// Device 推广设备
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:计算机
	// 1:移动设备
	// 取值与旧版报告文档不一致，请注意区分。
	Device int `json:"device,omitempty"`
	// MarketingTargetEnum 营销目标
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:网站链接
	// 1:应用推广
	// 2:门店推广
	// 3:推广营销活动
	// 4:电商店铺推广
	// 5:商品目录
	MarketingTargetEnum int `json:"marketingTargetEnum,omitempty"`
	// TargetingType 购买方式
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:关键词
	// 6:网址定向
	TargetingType int `json:"targetingType,omitempty"`
	// WinfoIdTypeEnum 定向类型
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:关键词
	// 1:词包
	// 3:自动扩量
	WinfoIdTypeEnum int `json:"winfoIdTypeEnum,omitempty"`
	// WInfoNameStatus 关键词/网址
	// 报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"关键词/网址[已删除]"。
	// 未删除物料正常显示名称，没有额外标记。
	WInfoNameStatus string `json:"wInfoNameStatus,omitempty"`
	// WInfoId 关键词ID
	WInfoId uint64 `json:"wInfoId,omitempty"`
	// MixWmatchEnum 匹配模式
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:智能匹配
	// 127:分匹配出价
	// 16:智能匹配核心词
	// 17:短语匹配
	// 48:精确匹配
	MixWmatchEnum int `json:"mixWmatchEnum,omitempty"`
	// QualityEnum 质量度
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 1:1
	// 10:10
	// 2:2
	// 3:3
	// 4:4
	// 5:5
	// 6:6
	// 7:7
	// 8:8
	// 9:9
	QualityEnum int `json:"qualityEnum,omitempty"`
	// EstimateClickRate 预估点击率
	// 用户看到您的关键词创意后，点击广告的可能性
	// 0: 数据积累中
	// 1: 低于平均
	// 2: 平均水平
	// 3: 高于平均
	EstimateClickRate int `json:"estimateClickRate,omitempty"`
	// BusinessRelationship 创意相关性
	// 您所使用的关键词和创意，与用户搜索意图之间的相关性
	// 0: 数据积累中
	// 1: 低于平均
	// 2: 平均水平
	// 3: 高于平均
	BusinessRelationship int `json:"businessRelationship,omitempty"`
	// LandPageExperience 落地页体验
	// 用户点击您的关键词广告后，前往广告落地页所感受到的用户体验和质量
	// 0: 数据积累中
	// 1: 低于平均
	// 2: 平均水平
	// 3: 高于平均
	LandPageExperience int `json:"landPageExperience,omitempty"`
	// WMatchId 触发模式
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 101:自动定向（智能匹配）
	// 103:网址投放
	// 109:商品定向
	// 110:自动定向
	// 111:词包
	// 15:智能
	// 16:智能匹配-人群智选
	// 31:短语
	// 63:精确
	// 该指标不支持小时。不能与provinceCityName同时使用。
	WMatchId int `json:"wMatchId,omitempty"`
	// QueryWord 搜索词
	QueryWord string `json:"queryWord,omitempty"`
	// QueryStatusName 账户添加状态
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 0:已添加
	// 1:未添加
	// 2:不可添加
	QueryStatusName int `json:"queryStatusName,omitempty"`
	// OuterId 商品ID（商家内部ID）
	OuterId string `json:"outerId,omitempty"`
	// ProductNameStatus 商品名称
	// 报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"商品名称[已删除]"。
	// 未删除物料正常显示名称，没有额外标记。
	ProductNameStatus string `json:"productNameStatus,omitempty"`
	// AudienceTagId 人群编号
	AudienceTagId uint64 `json:"audienceTagId,omitempty"`
	// AudienceName 人群名称
	AudienceName string `json:"audienceName,omitempty"`
	// AudienceStatusEnum 类型
	// 查询时使用key，数据返回value.
	// 枚举值：
	// 0:基本属性
	// 1:自定义人群
	// 1025:已转化用户
	// 2:设备属性
	// 4:商圈地域
	AudienceStatusEnum int `json:"audienceStatusEnum,omitempty"`
	// SegmentId SegmentId
	// 10101: 图集，10102: 大图，10104: 单图，10105:图文
	SegmentId int `json:"segmentId,omitempty"`
	// BindId 标识一次绑定
	BindId uint64 `json:"bindId,omitempty"`
	// VideoId 视频ID
	VideoId uint64 `json:"videoId,omitempty"`
	// VideoNameStatus 视频名称
	// 报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"视频名称[已删除]"。
	// 未删除物料正常显示名称，没有额外标记。
	VideoNameStatus string `json:"videoNameStatus,omitempty"`
	// VideoUrl 视频URL
	VideoUrl string `json:"videoUrl,omitempty"`
	// VideoInfo 视频内容预览
	// 组合列，包含如下属性：
	// videoUrl：视频内容预览
	// videoThumbnailUrl：视频封面对应url
	// 组合列在返回数据时会放到一个内嵌的Map结构中，比如：
	// {"videoInfo":{"videoThumbnailUrl":"xxx","videoUrl":"xxx"}}
	VideoInfo *VideoInfo `json:"videoInfo,omitempty"`
	// PicUrl 图片URL
	PicUrl string `json:"picUrl,omitempty"`
	// SegmentSign 图片签名
	SegmentSign string `json:"segmentSign,omitempty"`
	// SegmentType 比例
	// 只能取以下值： 101 1:1
	// 100 1:1
	// 320 3:1
	// 321 1.77:1
	SegmentType int `json:"segmentType,omitempty"`
	// SegmentDesc 主题
	SegmentDesc string `json:"segmentDesc,omitempty"`
	// AdgroupNameBind 绑定范围
	AdgroupNameBind string `json:"adgroupNameBind,omitempty"`
	// PicScale 图片比例
	// 查询时使用key，数据返回value.
	// 枚举值：
	// 1: 1:1
	// 2: 1.61:1
	// 3: 3:1
	// 4: 1.77:1
	PicScale int `json:"picScale,omitempty"`
	// ClickRegion 点击区域
	// 仅支持用于filters中进行数据筛选
	// 0为仅限图片，不加过滤为整体广告
	ClickRegion int `json:"clickRegion,omitempty"`
	// ProductSegment 组件类型
	// 查询时使用key，数据返回value。
	// 枚举值：
	// 1040:交通票务列表
	// 11006:全息图文
	// 11007:商品橱窗
	// 11009:商品精选
	// 11014:全息
	// 11015:商品列表
	// 11021:单品橱窗
	// 11023:图文商品列表
	// 11026:单品爆款橱窗
	// 11028:横划大卡片橱窗
	// 11029:文本横划
	// 11031:智选列表
	// 11033:结构化列表
	// 11034:全息服务直达
	// 11035:智选列表服务直达
	// 11037:站内直达高级版
	// 11038:知识直达
	// 11041:全息摘要
	// 11043:单品多图
	// 11044:商品卡片精选
	// 11045:商品卡片
	// 11052:单品大图
	// 11055:活动大图
	// 11058:类目橱窗
	// 11059:搜索推广头条
	// 11061:导航
	// 11063:站内直达
	// 11069:横图卡片
	// 11072:百科服务化
	// 3037:全景橱窗
	// 3198:热点橱窗
	ProductSegment int `json:"productSegment,omitempty"`
	// ProviceName 省
	// 该指标不支持小时。
	ProviceName string `json:"proviceName,omitempty"`
	// ProviceCityName 地市
	// 该指标不支持小时。
	ProviceCityName string `json:"proviceCityName,omitempty"`
}

// Metric 指标
type Metric struct {
	// BidNew 出价
	BidNew float64 `json:"bidNew,omitempty"`
	// TransPrice 目标转化成本
	TransPrice float64 `json:"transPrice,omitempty"`
	// Impression 表示展现次数。
	Impression int64 `json:"impression,omitempty"`
	// Click 表示点击次数。
	Click int64 `json:"click,omitempty"`
	// Cost 表示消费金额。
	Cost float64 `json:"cost,omitempty"`
	// CTR 表示点击率。 计算公式：Click/Impression。
	CTR float64 `json:"ctr,omitempty"`
	// CPC 表示平均点击价格。 计算公式：Cost/Click。
	CPC float64 `json:"cpc,omitempty"`
	// CPM 表示千次展现消费金额。 计算公式：Cost/(Impression/1000)。
	CPM float64 `json:"cpm,omitempty"`
	// TopPageViews 上方位展现
	TopPageViews int64 `json:"topPageViews,omitempty"`
	// TopPClicks 上方位点击
	TopPClicks int64 `json:"topPClicks,omitempty"`
	// TopPay 上方位消费
	TopPay int64 `json:"topPay,omitempty"`
	// TopPvWinA 上方展现胜出率
	// 【上方展现胜出率】与【推广业务】互斥并该指标不能与provinceCityName,cityName同时请求！
	// 上方展现胜出率=账户上方展现／账户符合投放条件的搜索次数
	TopPvWinA float64 `json:"topPvWinA,omitempty"`
	// TopPvWinP 上方展现胜出率（精确)
	// 【上方展现胜出率（精确）】与【推广业务】互斥 并该指标不能与provinceCityName,cityName同时请求！
	// 上方展现胜出率（精确）=账户精确触发的上方展现／账户符合投放条件&符合精确触发的搜索次数
	TopPvWinP float64 `json:"topPvWinP,omitempty"`
	// TopFirstPageViews 上方位首位展现
	TopFirstPageViews int64 `json:"topFirstPageViews,omitempty"`
	// TopFirstPvWinA 上方首位胜出率
	// 【上方首位胜出率】与【推广业务】互斥
	TopFirstPvWinA float64 `json:"topFirstPvWinA,omitempty"`
	// TopFirstWinP 上方首位胜出率（精确)
	// 【上方首位胜出率（精确）】与【推广业务】互斥
	TopFirstWinP float64 `json:"topFirstWinP,omitempty"`
	// PhoneButtonClicks 表示组件点击次数。 用于点击创意组件 button 次数。
	PhoneButtonClicks int64 `json:"phoneButtonClicks,omitempty"`
	// Interaction 表示广告的互动次数。 对于视频广告，互动行为包含点击互动和有效播放互动，对其他格式的广告，互动依然仅包含点击。
	Interaction int64 `json:"interaction,omitempty"`
}

// IdeaInfo 创意
type IdeaInfo struct {
	// IdeaId：创意ID
	IdeaId uint64 `json:"ideaId,omitempty"`
	// IdeaTitle：创意标题
	IdeaTitle string `json:"ideaTitle,omitempty"`
	// IdeaTitleStatus：创意标题
	IdeaTitleStatus string `json:"ideaTitleStatus,omitempty"`
	// IdeaShowURL：显示URL
	IdeaShowURL string `json:"ideaShowURL,omitempty"`
	// IdeaMShowURL：显示URL(移动)
	IdeaMShowURL string `json:"ideaMShowURL,omitempty"`
	// IdeaDesc：创意描述1
	IdeaDesc string `json:"ideaDesc,omitempty"`
	// IdeaDesc2：创意描述2
	IdeaDesc2 string `json:"ideaDesc2,omitempty"`
}

// VideoInfo 视频内容预览
type VideoInfo struct {
	// VideoUrl：视频内容预览
	VideoUrl string `json:"videoUrl,omitempty"`
	// VideoThumbnailUrl：视频封面对应url
	VideoThumbnailUrl string `json:"videoThumbnailUrl,omitempty"`
}
