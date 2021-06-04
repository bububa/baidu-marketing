package balance

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetPaymentRecordRequest struct {
	FundType  int                         `json:"fundType,omitempty"`  // 查询付款记录类型，不同类型返回的参数不同，详见返回结果：1：直销客户实际资金流动; 2：直销客户优惠资金流动;21：KA实际资金流动;22：KA优惠资金流动;23：KA待加款记录
	condition *model.Condition            `json:"condition,omitempty"` // 查询条件对象
	ChunkSize int64                       `json:"chunkSize,omitempty"` // 单次请求数据行数，默认500
	Offset    map[string]map[string]int64 `json:"offset,omitempty"`    // 返回数据行数偏移值，只支持id的gt（大于）、lt（小于），如：{ "id": { "gt": 0 } }
	Sort      map[string][]string         // 排序字段和排序方向列表，只支持id的asc或desc，如：{"in": [ "id asc",  ]}
}

func (r GetPaymentRecordRequest) Url() string {
	return fmt.Sprintf("%sPaymentService/getPaymentRecord", model.BASE_URL_SMS)
}
