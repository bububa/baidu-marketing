package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type AddAdgroupRequest struct {
	AdgroupTypes []Adgroup `json:"adgroupTypes"`     // 新增推广单元物料;集合长度限制：[1, 5000]
	AdType       int       `json:"adType,omitempty"` // 投放广告类型;0 - 普通单元;14 - 商品单元;6 - 网址定向单元
}

func (r AddAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/addAdgroup", model.BASE_URL_SMS)
}
