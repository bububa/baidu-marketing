package account

// RTASetting rta 设置
type RTASetting struct {
	// UserId 账户ID
	UserId int64 `json:"userId,omitempty"`
	// Url 接口地址，长度不超过1024个字符，http/https开头，不含空格，必填
	Url string `json:"url,omitempty"`
	// Qps 请求qps，大于等于0，0表示不限制（默认），单位次/秒，选填，此字段暂不生效
	Qps *int `json:"qps,omitempty"`
	// CacheTime 缓存时长，大于等于0，0表示不使用缓存（默认），单位秒，选填，此字段暂不生效
	CacheTime *int64 `json:"cacheTime,omitempty"`
	// StrategyIds 策略ids，大于0，上限100个，选填
	StrategyIds []int64 `json:"strategyIds,omitempty"`
	// CustomToken 自定义token，最大256个字符，选填
	CustomToken string `json:"customToken,omitempty"`
	// Token 百度认证token，系统自动生成，新增无效
	Token string `json:"token,omitempty"`
}
