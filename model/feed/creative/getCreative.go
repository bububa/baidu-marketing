package creative

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetCreativeRequest 查询创意 API Request
type GetCreativeRequest struct {
	// Ids 查询id集合;类型为单元ID时不超过100个，类型为创意ID时不超过100个，建议分批多次请求
	Ids []uint64 `json:"ids,omitempty"`
	// CreativeFeedFields 查询推广创意字段
	// 取值范围：枚举值，列表如下
	// creativeFeedId - 创意id
	// adgroupFeedId - 推广单元id
	// materialstyle - 样式
	// creativeFeedName - 创意名称
	// pause - 暂停/启用创意
	// material - 物料内容
	// status - 创意状态
	// refusereason - 拒绝原因
	// expmask - 产品标示位
	// changeorder - 三图换顺序开关
	// commentnum - 显示评论数开关
	// readnum - 显示阅读数开关
	// playnum - 显示播放数开关
	// ideaType - 创意类型
	// showMt - 程序化创意展示样式
	// addtime - 创意添加时间
	// progFlag - 程序化创意工具标识
	// approvemsgnew - 格式化的拒绝理由
	// auditTimeModel - 产生类型
	// template - 商品目录创意参数
	// huitus - 商品目录创意绑定慧图模板id
	CreativeFeedFields []string `json:"creativeFeedFields,omitempty"`
	// IdType 查询id类型
	// 取值范围：枚举值，列表如下
	// 1 - 计划ID
	// 2 - 单元ID
	// 3 - 创意ID
	IdType int `json:"idType"`
	// CreativeFeedFilter 查询创意的筛选条件
	CreativeFeedFilter *CreativeFeedFilter `json:"creativeFeedFilter"`
}

// CreativeFeeeFilter 查询创意的筛选条件
type CreativeFeedFilter struct {
	// Status 创意状态过滤
	// 取值范围：枚举值，列表如下
	// 0 - 有效
	// 1 - 暂停推广
	// 2 - 审核中
	// 3 - 审核未通过
	// 5 - 部分有效
	// 6 - 未审核
	Status []int `json:"status"`
}

func (r GetCreativeRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "CreativeFeedService/getCreativeFeed")
}

// GetCreativeResponse 查询创意 API Response
type GetCreativeResponse struct {
	Data []Creative `json:"data,omitempty"`
}
