package balance

import "encoding/json"

type GetAccountTransferHistoryResponse struct {
	Data []AccountTransferHistory `json:"data,omitempty"`
}

type AccountTransferHistory struct {
	Optid         int64       `json:"optid,omitempty"`         // 操作人id
	OptName       string      `json:"optName,omitempty"`       // 操作人名称
	OutUcId       int64       `json:"outUcId,omitempty"`       // 转出账户id
	OutUcName     string      `json:"outUcName,omitempty"`     // 转出账户名称
	InUcId        int64       `json:"inUcId,omitempty"`        // 转入账户id
	InUcName      string      `json:"inUcName,omitempty"`      // 转入账户名称
	TransferType  string      `json:"transferType,omitempty"`  // 转账类型
	TransferMoney json.Number `json:"transferMoney,omitempty"` // 转账金额, 单位：元
	Uuid          string      `json:"uuid,omitempty"`          // 转账记录id
	TransferTime  string      `json:"transferTime,omitempty"`  // 转账时间
}
