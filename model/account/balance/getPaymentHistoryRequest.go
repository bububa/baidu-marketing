package balance

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetPaymentHistoryRequest 查询待加款信息
type GetPaymentHistoryRequest struct {
	// FundType 资金类型，必填，21（ka 现金），22（ka 优惠）
	FundType int `json:"fundType,omitempty"`
	// Id 记录id，第一页数据不填，第二页及往后的数据必须填，升序填当前页最大id值，降序填当前页最小id值，当前只支持降序
	Id int64 `json:"id,omitempty"`
	// PageSize 每页数量大小，必填，每页数量必须大于0，小于等于1000
	PageSize int64 `json:"pageSize,omitempty"`
	// StartTime 支付时间-开始时间，选填，格式：yyyy-MM-dd HH:mm:ss
	StartTime string `json:"startTime,omitempty"`
	// EndTime 支付时间-结束时间，选填，格式：yyyy-MM-dd HH:mm:ss
	EndTime string `json:"endTime,omitempty"`
}

func (r GetPaymentHistoryRequest) Url() string {
	return fmt.Sprintf("%sPaymentFeedService/getPaymentHistory", model.BASE_URL_FEED)
}
