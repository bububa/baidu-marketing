package keyword

import "github.com/xiaoshouchen/baidu-marketing/model"

// Keyword 关键词对象
type Keyword struct {
	// CampaignId 推广计划ID
	CampaignId uint64 `json:"campaign_id,omitempty"`
	// KeywordId 关键词ID
	KeywordId uint64 `json:"keyword_id,omitempty"`
	// AdgroupId 推广单元ID
	AdgroupId uint64 `json:"adgroup_id,omitempty"`
	// Keyword 关键词字面
	// 长度限制：最大40个字节，1个中文按2个字节计算，英文、数字按1个字节计算；
	// 修改时，仅支持在匹配模式为"智能匹配-核心词"时，修改{}圈选的核心词范围，不支持字面修改；
	// 如“鲜花配送”可改为“{鲜花}配送；
	// 营销目标为“商品目录”时可在两个#之间写商品中心的商品属性来新建模板词，模板词中的商品属性会自动替换为所有绑定的商品的对应属性来批量生成关键词，如#name#配送
	Keyword string `json:"keyword,omitempty"`
	// Price 关键词竞价价格
	// 取值范围：(0, 999.99]
	// 需要小于等于所属计划预算。当使用单元出价时，查询该字段将不会返回；
	// 修改时设置0表示取消该关键词出价，此时采用所属单元出价
	Price *float64 `json:"price,omitempty"`
	// MatchType 匹配模式
	// 与phraseType配合使用，列表如下:
	// matchType=1且phraseType=1：精确匹配；
	// matchType=2且phraseType=1：短语匹配；
	// matchType=2且phraseType=3：智能匹配
	MatchType int `json:"match_type,omitempty"`
	// Pause 启用/暂停关键词
	// 默认值：false
	// 取值范围：枚举值，列表如下
	// true - 暂停
	// false - 启用
	Pause *bool `json:"pause,omitempty"`
	// Status 关键词状态
	// 取值范围：枚举值，列表如下
	// 40 - 有效-移动url审核中
	// 41 - 有效
	// 42 - 暂停推广
	// 43 - 不宜推广
	// 44 - 展现受限：搜索无效
	// 45 - 待激活
	// 46 - 审核中
	// 47 - 搜索量过低
	// 48 - 部分无效
	// 49 - 展现受限：计算机搜索无效
	// 50 - 展现受限：移动搜索无效
	// 58 - 展现受限
	// 66 - 未审核
	// 备注：部分无效：投放设备为全部设备时，移动物料审核未过，计算机物料审核通过时，显示的是“部分无效”；
	// 搜索无效：当计算机出价低于计算机最低展现价格，则显示“计算机搜索无效”，当移动出价低于移动最低展现价格，则显示“移动搜索无效”，都无效时，显示搜索无效
	Status int `json:"status,omitempty"`
	// PcDestinationUrl 计算机访问网址
	// 域名需和账户网站URL域名相同，长度限制：最大1024字节;
	// 修改时设置为""（空字符串），表示取消URL设置;
	// 营销目标为“本地推广”下不支持该字段
	PcDestinationUrl *string `json:"pc_destination_url,omitempty"`
	// ModileDestinationUrl 移动访问网址
	// 域名需和账户网站URL域名相同，长度限制：最大1024字节；
	// 修改时设置为""（空字符串），表示取消URL设置;
	// 营销目标为“本地推广”下不支持该字段
	ModileDestinationUrl *string `json:"modile_destination_url,omitempty"`
	// PhraseType 细分匹配模式
	// 与matchType配合使用，列表如下:
	// matchType=1且phraseType=1：精确匹配；
	// matchType=2且phraseType=1：短语匹配；
	// matchType=2且phraseType=3：智能匹配
	PhraseType int `json:"phrase_type,omitempty"`
	// OfflineReason 物料推广下线原因
	// 当有多个推广下线原因时，数组会有多个元素，每个代表一种原因
	OfflineReason []model.OfflineReason `json:"offline_reason,omitempty"`
	// Tabs 关键词物料标签数组
	// 取值范围：[0, 31]
	// 说明：同一关键词支持添加多个标签，最多添加31个。
	// 举例：关键词“鲜花“同时标记了1、2、3个标签，则用[1,2,3]表示
	// 特殊说明：0表示无标签，31表示重点关键词标签
	Tabs []int `json:"tabs,omitempty"`
	// LeftPriceGuide 计算机指导价
	// 取值范围：[0, 999.99)
	// 说明：当历史数据不足时会返回-
	LeftPriceGuide *float64 `json:"left_price_guide,omitempty"`
	// MPriceGuide 移动指导价
	// 取值范围：[0, 999.99)
	// 说明：当历史数据不足时会返回-
	MPriceGuide *float64 `json:"m_price_guide,omitempty"`
	// Deeplink 应用调起网址
	// 长度限制：[0, 2048]
	// 该字段属于小流量功能，需要使用请向营销顾问申请（不在小流量名单中调用将报错）
	Deeplink *string `json:"deeplink,omitempty"`
	// MiniProgramUrl 小程序访问网址
	// 长度不超过1024字节
	MiniProgramUrl *string `json:"mini_program_url,omitempty"`
	// Quality 质量度
	// 取值范围：[0, 10]
	// 只读，质量度得分越高，越可能以较低的价格获得较为理想的推广结果排名。
	// 质量度得分1-2分是展现受限，即使有展现您所付出的点击成本也远高于行业评价水平
	// 完整质量度介绍建议参考：
	// https://yingxiao.baidu.com/course/outside/detail?from=yingxiaoSearch&detailId=3816
	Quality int `json:"quality,omitempty"`
	// EstimatedClickRate 预估点击率
	// 取值范围：枚举值，列表如下
	// 0 - 数据积累中
	// 1 - 低于平均
	// 2 - 平均水平
	// 3 - 高于平均
	EstimatedClickRate int `json:"estimated_click_rate,omitempty"`
	// BusinessRelationship 创意相关性
	// 取值范围：枚举值，列表如下
	// 0 - 数据积累中
	// 1 - 低于平均
	// 2 - 平均水平
	// 3 - 高于平均
	BusinessRelationship int `json:"business_relationship,omitempty"`
	// LandpageExperience 落地页体验
	// 取值范围：枚举值，列表如下
	// 0 - 数据积累中
	// 1 - 低于平均
	// 2 - 平均水平
	// 3 - 高于平均
	LandpageExperience int `json:"landpage_experience,omitempty"`
	// CreateTime 关键词创建时间
	CreateTime string `json:"create_time,omitempty"`
	// Ulink ios应用调起
	// 长度限制4096字符
	// 该字段属于小流量功能，需要使用请向营销顾问申请（不在小流量名单中调用将报错）
	// 当前仅支持营销目标-网站链接、应用推广、电商店铺（非基木鱼、非健康商城交易平台）的关键词层级
	Ulink *string `json:"ulink,omitempty"`
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
