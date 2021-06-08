package creative

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type DeleteCreativeRequest struct {
	CreativeIds []int64 `json:"creativeIds"`
}

func (r DeleteCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeService/deleteCreative", model.BASE_URL_SMS)
}
