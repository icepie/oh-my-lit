package jw

// JwTime 教务时间
type JwTime struct {
	Year       uint   `json:"year"`
	Term       uint   `json:"term"`
	Week       uint   `json:"week"`
	IsVacation bool   `json:"is_vacation"`
	RawData    string `json:"raw_data"`
	OnlineNum  uint   `json:"online_num"`
	Sub        string `json:"sub"`
}

// 课程信息
type CourseInfo struct {
	Day      int    `json:"day"`
	Titile   string `json:"title"`
	Location string `json:"location"`
	Credit   string `json:"credit"`
	Code     string `json:"code"`
	Sections []int  `json:"sections"`
	Start    int    `json:"start"`
	Duration int    `json:"duration"`
	Teacher  string `json:"teacher"`
	Weeks    []int  `json:"weeks"`
	Time     string `json:"time"`
}
