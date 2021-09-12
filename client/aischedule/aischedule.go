package aischedule

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	AISchedulParserURL = "https://open-schedule.ai.xiaomi.com/api/schedule/parser"
	// AIScheduleTBID        = "45625"
	AIScheduleTableURL    = "https://i.ai.mi.com/course-multi/table"
	AIScheduleTablesURL   = "https://i.ai.mi.com/course-multi/tables"
	AIScheduleFeedbackURL = "https://open-schedule.ai.xiaomi.com/api/feedback"
	AIScheduleCourseURL   = "https://i.ai.mi.com/course-multi/courseInfos"
	AIScheduleShareURL    = "https://i.ai.mi.com/course-multi/shareToken"
	AIScheduleImportURL   = "https://i.ai.mi.com/h5/precache/ai-schedule"

	// 默认课表风格
	StyleList = []ScheduleStyle{{Color: "#00A6F2", Background: "#E5F4FF"},
		{Color: "#FC6B50", Background: "#FDEBDE"},
		{Color: "#3CB3C8", Background: "#DEFBF8"},
		{Color: "#7D7AEA", Background: "#EDEDFF"},
		{Color: "#FF9900", Background: "#FCEBCD"},
		{Color: "#EF5B75", Background: "#FFEFF0"},
		{Color: "#5B8EFF", Background: "#EAF1FF"},
		{Color: "#F067BB", Background: "#FFEDF8"},
		{Color: "#29BBAA", Background: "#E2F8F3"},
		{Color: "#CBA713", Background: "#FFF8C8"},
		{Color: "#B967E3", Background: "#F9EDFF"},
		{Color: "#6E8ADA", Background: "#F3F2FD"},
		{Color: "#00A6F2", Background: "#E5F4FF"},
		{Color: "#FC6B50", Background: "#FDEBDE"},
		{Color: "#3CB3C8", Background: "#DEFBF8"}}
)

type AIScheduleUser struct {
	UserID   int64
	DeviceID string // 最重要
	Client   *resty.Client
}

// NewHealthUser 新建健康平台用户
func NewAISchedule() *AIScheduleUser {

	var u AIScheduleUser

	u.Client = resty.New()

	// u.Client.SetDebug(true)
	// u.Client.SetHeaders(MainHeaders)
	u.Client.SetTimeout(5 * time.Second)

	return &u
}

// SetUserID 设置用户ID
func (u *AIScheduleUser) SetUserID(userID int64) *AIScheduleUser {
	u.UserID = userID
	return u
}

// SetDeviceID 设置密码
func (u *AIScheduleUser) SetDeviceID(deviceID string) *AIScheduleUser {
	u.DeviceID = deviceID
	return u
}
