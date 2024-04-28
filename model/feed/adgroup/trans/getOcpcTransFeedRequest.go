package trans

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetOcpcTransFeedRequest 查询转化追踪 API Request
type GetOcpcTransFeedRequest struct {
	// 指定查询的接入方式
	// 取值范围：枚举值，列表如下
	// 0 - 包含全部接入方式
	// 1 - 应用API
	// 2 - 基木鱼营销页
	// 4 - API激活
	// 5 - 网页JS布码
	// 7 - 线索API
	// 8 - 咨询工具授权
	// 9 - 百度智能小程序
	// 13 - 应用SDK
	// 23 - 百度统计网站导入
	// 24 - 百度统计小程序导入
	// 28 - 百度小说书城
	TransForm int `json:"transFrom,omitempty"`
	// JmyPageFilter 过滤条件 当transFrom=2时选填，可以过滤获取基木鱼营销页的数据
	JmyPageFilter *JmyPageFilter `json:"jmyPageFilter,omitempty"`
}

// JmyPageFilter 过滤条件
type JmyPageFilter struct {
	// ShowType 落地页类型
	// 取值范围：枚举值，列表如下
	// 0 - h5
	// 1 - pc
	// 2 - 小程序
	ShowType int `json:"showType"`
	// PlatformIds 平台
	//     取值范围：枚举值，列表如下
	// 1 - 基木鱼平台
	// 2 - 基木鱼电商
	PlatformIds []int `json:"platformIds"`
	// SearchFields 搜索字段
	SearchFields *SearchFieldType `json:"searchFields"`
}

// SearchFieldType 搜索字段
type SearchFieldType struct {
	// PageName 页面名称
	// 取值范围：[1, 50]
	// 1个中文按2个字符计算，英文、数字按1个字符计算
	PageName string `json:"pageName,omitempty"`
	// Id 通过id 搜索
	// 基木鱼平台平台通过 pageId 搜索，基木鱼电商平台通过 spuId 搜索
	Id uint64 `json:"id,omitempty"`
	// SearchKey 用户搜索框输入内容
	// 非空searchKey字段会把该值赋值给pageName,同时如果该值字面量是数值该值还会被赋值给id
	SearchKey string `json:"searchKey,omitempty"`
}

func (r GetOcpcTransFeedRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "SearchFeedService/getOcpcTransFeed")
}

// GetOcpcTransFeedResponse 查询转化追踪 API Response
type GetOcpcTransFeedResponse struct {
	Data []OcpcTransFeed `json:"data,omitempty"`
}
