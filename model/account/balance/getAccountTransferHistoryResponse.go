package balance

import "encoding/json"

// GetAccountTransferHistoryResponse 查询转账记录 API Response
type GetAccountTransferHistoryResponse struct {
	Data []AccountTransferHistory `json:"data,omitempty"`
}

// AccountTransferHistory 账户转账记录
type AccountTransferHistory struct {
	// Optid 操作人id
	Optid int64 `json:"optid,omitempty"`
	// OptName 操作人名称
	OptName string `json:"optName,omitempty"`
	// OutUcId 转出账户id
	OutUcId int64 `json:"outUcId,omitempty"`
	// OutUcName 转出账户名称
	OutUcName string `json:"outUcName,omitempty"`
	// InUcId 转入账户id
	InUcId int64 `json:"inUcId,omitempty"`
	// InUcName 转入账户名称
	InUcName string `json:"inUcName,omitempty"`
	// TransferType 转账类型
	TransferType string `json:"transferType,omitempty"`
	// TransferMoney 转账金额, 单位：元
	TransferMoney json.Number `json:"transferMoney,omitempty"`
	// Uuid 转账记录id
	Uuid string `json:"uuid,omitempty"`
	// TransferTime 转账时间
	TransferTime string `json:"transferTime,omitempty"`
}
