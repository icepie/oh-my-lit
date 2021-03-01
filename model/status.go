package model

// JwTime 教务时间
type JwTime struct {
	Date string      `json:"date"`
	Day  string      `json:"day"`
	Term string      `json:"term"`
	Week interface{} `json:"week"`
}

// JWStatus 教务在线结构
type JWStatus struct {
	OnlineNumber interface{} `json:"online_number"`
	JwTime       string      `json:"jw_time"`
}
