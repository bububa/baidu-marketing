package newimage

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

// GetImageRequest （新图片服务（ImageManageService）打通了搜索图库与信息流图库）获取图片素材API Request
type GetImageRequest struct {
	// Filters 过滤条件
	Filters []FilterCondition `json:"filters,omitempty"`
	// Page 当前页；选填，默认为1
	Page int `json:"pageNo,omitempty"`
	// PageSize 当前页大小；选填 ， 默认40，范围（0, 1000)
	PageSize int `json:"pageSize,omitempty"`
}

func (r GetImageRequest) Url() string {
	return fmt.Sprintf("%sImageManageService/getImageList", model.BASE_URL_SMS)
}

// FilterCondition 过滤条件
// OP 取值范围：枚举值。
// eq：等于；
// in：包含；
// ge：大于等于；
// gt：大于；
// le：小于等于；
// lt：小于；
// bt：在...范围
type FilterCondition struct {
	// Field 支持过滤的字段
	Field string `json:"field,omitempty"`
	// OP 操作符
	OP string `json:"op,omitempty"`
}
