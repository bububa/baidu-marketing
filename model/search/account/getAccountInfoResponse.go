package account

// GetAccountInfoResponse 查询账户 API Response
type GetAccountInfoResponse struct {
	Data []Account `json:"data,omitempty"`
}
