package jw

// JwTime 教务时间
type JwTime struct {
	Year       uint   `json:"year"`
	Term       uint   `json:"term"`
	Week       uint   `json:"week"`
	IsVacation bool   `json:"is_vacation"`
	RawData    string `json:"raw_data"`
	OnlineNum  uint   `json:"online_num"`
}
