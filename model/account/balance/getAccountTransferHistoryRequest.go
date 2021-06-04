package balance

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetAccountTransferHistoryRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (r GetAccountTransferHistoryRequest) Url() string {
	return fmt.Sprintf("%sMccHistoryService/getAccountTransferHistory", model.BASE_URL_SMS)
}
