package model

// JwTime 教务时间
type JwTime struct {
	Term interface{} `json:"term"`
	Week interface{} `json:"week"`
}

// JWStatus 教务在线结构
type JWStatus struct {
	OnlineNumber interface{} `json:"online_number"`
	JwTime       JwTime      `json:"jw_time"`
}
