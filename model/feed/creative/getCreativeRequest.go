package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetCreativeRequest 查询推广创意字段
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

type GetCreativeRequest struct {
	// Ids 查询id集合;类型为单元ID时不超过100个，类型为创意ID时不超过100个，建议分批多次请求
	Ids []int64 `json:"ids"`
	// CreativeFeedFields 查询推广创意字段
	CreativeFeedFields []string `json:"creativeFeedFields"`
	// IdType 查询id类型;2 - 单元ID;3 - 创意ID
	IdType int `json:"idType"`
	// CreativeFeedFilter 查询创意的筛选条件
	CreativeFeedFilter *CreativeFeedFilter `json:"creativeFeedFilter"`
}

// CreativeFeeeFilter 查询创意的筛选条件
type CreativeFeedFilter struct {
	// Status 创意状态过滤
	Status []int `json:"status"`
}

func (r GetCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeFeedService/getCreativeFeed", model.BASE_URL_FEED)
}
