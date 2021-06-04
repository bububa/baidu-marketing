package model

type Condition struct {
	Paytime map[string]string `json:"paytime,omitempty"` // 加款时间范围，支持gte（大于等于）、lte（小于等于），如：{ "gte": "2016-11-22 00:00:00", "lte": "2016-11-23 00:00:00" }
	Status  map[string][]int  `json:"status,omitempty"`  // 要查询 ka 待加款流水支付状态，如{ "in": [0, 1, 2,]  }
}
