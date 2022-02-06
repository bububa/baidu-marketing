package account

// GetAccountFeedResponse 查询账户信息 API Response
type GetAccountFeedResponse struct {
	Data []Account `json:"data,omitempty"`
}
