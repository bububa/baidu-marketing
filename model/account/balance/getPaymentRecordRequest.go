package balance

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetPaymentRecordRequest 查询付款信息与待加款信息 API Request
type GetPaymentRecordRequest struct {
	// FundType 查询付款记录类型，不同类型返回的参数不同，详见返回结果：1：直销客户实际资金流动; 2：直销客户优惠资金流动;21：KA实际资金流动;22：KA优惠资金流动;23：KA待加款记录
	FundType int `json:"fundType,omitempty"`
	// Condition 查询条件对象
	Condition *Condition `json:"condition,omitempty"`
	// ChunkSize 单次请求数据行数，默认500
	ChunkSize int64 `json:"chunkSize,omitempty"`
	// Offset 返回数据行数偏移值，只支持id的gt（大于）、lt（小于），如：{ "id": { "gt": 0 } }
	Offset map[string]map[string]int64 `json:"offset,omitempty"`
	// Sort 排序字段和排序方向列表，只支持id的asc或desc，如：{"in": [ "id asc",  ]}
	Sort map[string][]string
}

func (r GetPaymentRecordRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "PaymentService/getPaymentRecord")
}

// GetPaymentRecordResponse 查询付款信息与待加款信息
type GetPaymentRecordResponse struct {
	// FundType 查询付款记录类型，不同类型返回的参数不同，详见返回结果：1：直销客户实际资金流动; 2：直销客户优惠资金流动;21：KA实际资金流动;22：KA优惠资金流动;23：KA待加款记录
	FundType int `json:"fundType,omitempty"`
	// Condition 查询条件对象
	Condition *Condition `json:"condition,omitempty"`
	// ChunkSize 单次请求数据行数，默认500
	ChunkSize int64 `json:"chunkSize,omitempty"`
	// Offset 返回数据行数偏移值，只支持id的gt（大于）、lt（小于），如：{ "id": { "gt": 0 } }
	Offset map[string]uint64 `json:"offset,omitempty"`
	// Sort 排序字段和排序方向列表，只支持id的asc或desc，如：{"in": [ "id asc",  ]}
	Sort []string
	// Total
	Total int64 `json:"total,omitempty"`
	// List
	List []PaymentRecord `json:"list,omitempty"`
}

// PaymentRecord
type PaymentRecord struct {
	// Id 流水号
	Id uint64 `json:"id,omitempty"`
	// Uid 用户id
	Uid uint64 `json:"uid,omitempty"`
	// Targetid 主账户id
	Targetid uint64 `json:"targetid,omitempty"`
	// Paytime 加款时间
	Paytime string `json:"paytime,omitempty"`
	// Rcvtime 到账时间
	Rcvtime string `json:"rcvtime,omitempty"`
	// PaymethodName 支付方式名称
	PaymethodName string `json:"paymethodName,omitempty"`
	// Fund 金额，单位：元，正数表示加款，负数表示退款
	Fund float64 `json:"fund,omitempty"`
	// ProductName 账户类型名称
	ProductName string `json:"productName,omitempty"`
	// Cash 现金金额
	Cash float64 `json:"cash,omitempty"`
	// Bonus 优惠金额
	Bonus float64 `json:"bonus,omitempty"`
	// CachetypeName 待加款类型
	CachetypeName string `json:"cachetypeName,omitempty"`
	// Actflag 支付状态id，枚举值: 0：待支付;1：已付款;2：已取消
	Actflag int `json:"actflag,omitempty"`
	// BillStatus 支付状态名称
	BillStatus string `json:"billStatus,omitempty"`
	// Orderrow 订单行号
	Orderrow uint64 `json:"orderrow,omitempty"`
}
