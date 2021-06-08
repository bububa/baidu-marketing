package adgroup

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type DeleteAdgroupRequest struct {
	AdgroupIds []int64 `json:"adgroupIds"`
}

func (r DeleteAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupService/deleteAdgroup", model.BASE_URL_SMS)
}
