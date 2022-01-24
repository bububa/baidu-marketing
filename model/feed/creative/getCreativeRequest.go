package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// 查询推广创意字段
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
	Ids                []int64  `json:"ids"`                // 查询id集合;类型为单元ID时不超过100个，类型为创意ID时不超过100个，建议分批多次请求
	CreativeFeedFields []string `json:"creativeFeedFields"` // 查询推广创意字段
	IdType             int      `json:"idType"`             // 查询id类型;2 - 单元ID;3 - 创意ID
	CreativeFeedFilter struct {
		Status []int `json:"status"`
	} `json:"creativeFeedFilter"`
}

func (r GetCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeFeedService/getCreativeFeed", model.BASE_URL_FEED)
}
