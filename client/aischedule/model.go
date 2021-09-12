package aischedule

// ScheduleStyle 课表风格
type ScheduleStyle struct {
	Color      string `json:"color"`
	Background string `json:"background"`
}

// Ret 返回结构
type Ret struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Desc string      `json:"desc"`
}

type CreateRet struct {
	Code int64  `json:"code"`
	Data int64  `json:"data"`
	Desc string `json:"desc"`
}

type ChangeRet struct {
	Code int64  `json:"code"`
	Data bool   `json:"data"`
	Desc string `json:"desc"`
}

type AddRet struct {
	Code int64   `json:"code"`
	Data []int64 `json:"data"`
	Desc string  `json:"desc"`
}

type ShareRet struct {
	Code int64 `json:"code"`
	Data struct {
		DueTime int64  `json:"dueTime"`
		Token   string `json:"token"`
	} `json:"data"`
	Desc string `json:"desc"`
}

type ParseRet struct {
	TBID      string `json:"tb_id"`
	ParserRet string `json:"parserRet"`
}

type GetCourseRet struct {
	Code int64 `json:"code"`
	Data struct {
		Courses []AppCourse `json:"courses"`
	} `json:"data"`
	Desc string `json:"desc"`
}

type TableSetting struct {
	AfternoonNum  int64  `json:"afternoonNum"`
	Extend        string `json:"extend"`
	ID            int64  `json:"id"`
	IsWeekend     int64  `json:"isWeekend"`
	MorningNum    int64  `json:"morningNum"`
	NightNum      int64  `json:"nightNum"`
	PresentWeek   int64  `json:"presentWeek"`
	School        string `json:"school"`
	Sections      string `json:"sections"`
	Speak         int64  `json:"speak"`
	StartSemester string `json:"startSemester"`
	TotalWeek     int64  `json:"totalWeek"`
	WeekStart     int64  `json:"weekStart"`
}

type AppCourse struct {
	AttendTime  int64         `json:"attendTime"`
	CreateTime  int64         `json:"createTime"`
	CtID        int64         `json:"ctId"`
	Day         int64         `json:"day"`
	Extend      string        `json:"extend"`
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Position    string        `json:"position"`
	SectionList []interface{} `json:"sectionList"`
	Sections    string        `json:"sections"`
	Style       string        `json:"style"`
	Teacher     string        `json:"teacher"`
	UpdateTime  int64         `json:"updateTime"`
	Weeks       string        `json:"weeks"`
}

type Table struct {
	Courses []interface{} `json:"courses"`
	Current int64         `json:"current"`
	ID      int64         `json:"id"`
	Name    string        `json:"name"`
	Setting struct {
		AfternoonNum  int64  `json:"afternoonNum"`
		Confirm       int64  `json:"confirm"`
		CreateTime    int64  `json:"createTime"`
		Extend        string `json:"extend"`
		ID            int64  `json:"id"`
		IsWeekend     int64  `json:"isWeekend"`
		MorningNum    int64  `json:"morningNum"`
		NightNum      int64  `json:"nightNum"`
		PresentWeek   int64  `json:"presentWeek"`
		School        string `json:"school"`
		SectionTimes  string `json:"sectionTimes"`
		Speak         int64  `json:"speak"`
		StartSemester string `json:"startSemester"`
		TotalWeek     int64  `json:"totalWeek"`
		UpdateTime    int64  `json:"updateTime"`
		WeekStart     int64  `json:"weekStart"`
	} `json:"setting"`
}

type Schedule struct {
	CourseInfos []struct {
		Day      int       `json:"day"`
		Name     string    `json:"name"`
		Position string    `json:"position"`
		Sections []Section `json:"sections"`
		Teacher  string    `json:"teacher"`
		Weeks    []int     `json:"weeks"`
	} `json:"courseInfos"`
	SectionTimes []struct {
		EndTime   string `json:"endTime"`
		Section   int    `json:"section"`
		StartTime string `json:"startTime"`
	} `json:"sectionTimes"`
}

type Section struct {
	Section int `json:"section"`
}

type Course struct {
	Day      int64  `json:"day"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Sections string `json:"sections"`
	Style    string `json:"style"`
	Teacher  string `json:"teacher"`
	Weeks    string `json:"weeks"`
}
