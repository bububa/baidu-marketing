package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetAccountLiveDataRequest 账户实时数据 API Request
type GetAccountLiveDataRequest struct {
	// DataType 账户/计划实时数据查询;选填，默认值：1;1：账户数据;2：计划数据
	DataType int `json:"dataType,omitempty"`
	// Device 投放设备; 选填，默认值：0 ; 0：全部设备; 1：仅计算机; 2：仅移动; 说明：仅在dataType = 2时有效
	Device int `json:"device,omitempty"`
}

func (r GetAccountLiveDataRequest) Url() string {
	return fmt.Sprintf("%sLiveReportService/getAccountLiveData", model.BASE_URL_SMS)
}
