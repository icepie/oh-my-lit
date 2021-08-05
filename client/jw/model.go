package jw

// JwTime 教务时间
type JwTime struct {
	Year       uint   `josn:"year"`
	Term       uint   `josn:"term"`
	Week       uint   `json:"week"`
	IsVacation bool   `josn:"is_vacation"`
	RawData    string `json:"raw_data"`
	OnlineNum  uint   `json:"online_num"`
}
