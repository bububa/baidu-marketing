package trans

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

type OcpcTransRequest struct {
	// 指定查询的接入方式
	TransForm int `json:"transFrom"`
	// 当transFrom=2时选填，可以过滤获取基木鱼营销页的数据
	JmyPageFilter JmyPageFilter `json:"jmyPageFilter"`
	// 搜索字段
	SearchFields SearchFieldType `json:"searchFields"`
}

func (r OcpcTransRequest) Url() string {
	return fmt.Sprintf("%sSearchFeedService/getOcpcTransFeed", model.BASE_URL_FEED)
}

type JmyPageFilter struct {
	// 落地页类型
	ShowType int `json:"showType"`
	// 平台
	PlatformIds []int `json:"platformIds"`
}

type SearchFieldType struct {
	// 页面名称
	PageName string `json:"pageName"`
	// 通过id 搜索
	Id int64 `json:"id"`
	// 用户搜索框输入内容
	SearchKey string `json:"searchKey"`
}
