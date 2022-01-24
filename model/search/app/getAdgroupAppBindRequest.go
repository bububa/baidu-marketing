package app

import (
	"fmt"
	"github.com/bububa/baidu-marketing/model"
)

type GetAdgroupAppBindRequest struct {
	IdType   int     `json:"idType,omitempty"`
	Ids      []int64 `json:"ids,omitempty"`
	Name     string  `json:"name,omitempty"`
	Platform []int   `json:"platform,omitempty"`
	Status   []int   `json:"status,omitempty"`
	OrderBy  string  `json:"orderBy,omitempty"`
	Desc     bool    `json:"desc,omitempty"`
	Limit    []int   `json:"limit,omitempty"`
}

func (r GetAdgroupAppBindRequest) Url() string {
	return fmt.Sprintf("%sAdgroupAppService/getAdgroupAppBind", model.BASE_URL_SMS)
}
