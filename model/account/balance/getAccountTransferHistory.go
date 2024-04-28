package balance

import (
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/util"
)

// GetAccountTransferHistoryRequest 查询转账记录 API Request
type GetAccountTransferHistoryRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (r GetAccountTransferHistoryRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "MccHistoryService/getAccountTransferHistory")
}

// GetAccountTransferHistoryResponse 查询转账记录 API Response
type GetAccountTransferHistoryResponse struct {
	Data []AccountTransferHistory `json:"data,omitempty"`
}

// AccountTransferHistory 账户转账记录
type AccountTransferHistory struct {
	// Optid 操作人id
	Optid uint64 `json:"optid,omitempty"`
	// OptName 操作人名称
	OptName string `json:"optName,omitempty"`
	// OutUcId 转出账户id
	OutUcId uint64 `json:"outUcId,omitempty"`
	// OutUcName 转出账户名称
	OutUcName string `json:"outUcName,omitempty"`
	// InUcId 转入账户id
	InUcId uint64 `json:"inUcId,omitempty"`
	// InUcName 转入账户名称
	InUcName string `json:"inUcName,omitempty"`
	// TransferType 转账类型
	TransferType string `json:"transferType,omitempty"`
	// TransferMoney 转账金额, 单位：元
	TransferMoney model.Float64 `json:"transferMoney,omitempty"`
	// Uuid 转账记录id
	Uuid string `json:"uuid,omitempty"`
	// TransferTime 转账时间
	TransferTime string `json:"transferTime,omitempty"`
}
