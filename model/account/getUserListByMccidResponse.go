package account

type GetUserListByMccidResponse struct {
	Data []MccUser `json:"data,omitempty"`
}

type MccUser struct {
	Userid   int64  `json:"userid,omitempty"`   // 用户id
	Username string `json:"username,omitempty"` // 用户名
	Mccid    int64  `json:"mccid,omitempty"`    // 账户管家id
	Fatname  string `json:"fatname,omitempty"`  // 账户管家名称
	Remark   string `json:"remark,omitempty"`   // 备注
}
