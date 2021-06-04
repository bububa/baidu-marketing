package balance

import "github.com/bububa/baidu-marketing/model"

type GetPaymentRecordResponse struct {
	FundType  int                         `json:"fundType,omitempty"`  // 查询付款记录类型，不同类型返回的参数不同，详见返回结果：1：直销客户实际资金流动; 2：直销客户优惠资金流动;21：KA实际资金流动;22：KA优惠资金流动;23：KA待加款记录
	condition *model.Condition            `json:"condition,omitempty"` // 查询条件对象
	ChunkSize int64                       `json:"chunkSize,omitempty"` // 单次请求数据行数，默认500
	Offset    map[string]map[string]int64 `json:"offset,omitempty"`    // 返回数据行数偏移值，只支持id的gt（大于）、lt（小于），如：{ "id": { "gt": 0 } }
	Sort      []string                    // 排序字段和排序方向列表，只支持id的asc或desc，如：{"in": [ "id asc",  ]}
	Total     int64                       `json:"total,omitempty"`
	List      []PaymentRecord             `json:"list,omitempty"`
}

type PaymentRecord struct {
	Id            int64   `json:"id,omitempty"`            // 流水号
	Uid           int64   `json:"uid,omitempty"`           // 用户id
	Targetid      int64   `json:"targetid,omitempty"`      // 主账户id
	Paytime       string  `json:"paytime,omitempty"`       // 加款时间
	Rcvtime       string  `json:"rcvtime,omitempty"`       // 到账时间
	PaymethodName string  `json:"paymethodName,omitempty"` // 支付方式名称
	Fund          float64 `json:"fund,omitempty"`          // 金额，单位：元，正数表示加款，负数表示退款
	ProductName   string  `json:"productName,omitempty"`   // 账户类型名称
	Cash          float64 `json:"cash,omitempty"`          // 现金金额
	Bonus         float64 `json:"bonus,omitempty"`         // 优惠金额
	CachetypeName string  `json:"cachetypeName,omitempty"` // 待加款类型
	Actflag       int     `json:"actflag,omitempty"`       // 支付状态id，枚举值: 0：待支付;1：已付款;2：已取消
	BillStatus    string  `json:"billStatus,omitempty"`    // 支付状态名称
	Orderrow      int64   `json:"orderrow,omitempty"`      // 订单行号
}
