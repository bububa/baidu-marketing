package keyword

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetWordRequest 查询关键词 API Request
type GetWordRequest struct {
	// WordFields 指定需要返回的关键词属性
	// 取值范围：枚举值，列表如下
	// keywordId - 关键词ID
	// campaignId - 推广计划ID
	// adgroupId - 推广单元ID
	// keyword - 关键词字面
	// price - 关键词竞价价格
	// pause - 启用/暂停关键词
	// matchType - 匹配模式
	// phraseType - 细分匹配模式
	// status - 关键词状态
	// apiInefficient - 是否低效关键词
	// pcDestinationUrl - 计算机访问网址
	// mobileDestinationUrl - 移动访问网址
	// tabs - 关键词物料标签数组
	// leftPriceGuide - 计算机指导价
	// mPriceGuide - 移动指导价
	// miniProgramUrl - 小程序访问网址
	// offlineReasons - 物料推广下线原因
	// createTime - 创建时间
	// ulink - ios调起应用（该功能目前在小流量，需要使用向营销顾问申请）
	// deeplink - 应用调起网址
	// quality - 质量度
	// estimatedClickRate - 预估点击率
	// businessRelationship - 创意相关性
	// landPageExperience - 落地页体验
	// pcFinalUrl - 计算机最终访问网址
	// pcTrackParam - 计算机监控后缀
	// pcTrackTemplate - 计算机第三方追踪模板
	// mobileFinalUrl - 移动最终访问网址
	// mobileTrackParam - 移动监控后缀
	// mobileTrackTemplate - 移动第三方追踪模板
	WordFields []string `json:"wordFields,omitempty"`
	// Ids 查询id集合
	// idType=5为单元ID，不超过50个；idType=11为关键词ID，不超过10000个
	Ids []uint64 `json:"ids,omitempty"`
	// IdType 查询id类型
	// 取值范围：枚举值，列表如下
	// 5 - 单元ID
	// 11 - 关键词ID
	IdType int `json:"idType,omitempty"`
	// GetTemp 是否查询关键词影子
	// 默认值：0
	// 0 - 只查询关键词本身，1 - 只查询关键词影子；
	// 想要获得关键词的全集，需要调用该方法两次，分别为getTemp=0和getTemp=1；
	// 影子说明：用户先向系统提交了关键词A，并且A已审核通过，之后再对A进行影响审核状态的修改（例如修改关键词url），修改后的关键词为A’（A’即为影子，仅对审核通过的物料进行修改才会产生影子），在A’通过审核生效之前，线上的生效创意仍然为A。 此时：getTemp为0查询到的是A，getTemp为1查询到的是A’
	GetTemp int `json:"getTemp,omitempty"`
}

func (r GetWordRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "KeywordService/getWord")
}

// GetWordResponse 查询关键词 API Response
type GetWordResponse struct {
	Data []Keyword `json:"data,omitempty"`
}
