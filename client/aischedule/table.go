package aischedule

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// GetTables 获取课程列表
func (u *AIScheduleUser) GetTables() (tables []Table, err error) {

	var result Ret

	_, err = u.Client.R().
		SetQueryParams(map[string]string{
			"userId":   strconv.FormatInt(u.UserID, 10),
			"deviceId": u.DeviceID,
		}).
		SetResult(&result).
		Get(AIScheduleTablesURL)

	if err != nil {
		return
	}

	if result.Code != 0 {
		err = errors.New(result.Desc)
		return
	}

	byteData, _ := json.Marshal(result.Data)
	err = json.Unmarshal(byteData, &tables)
	if err != nil {
		return
	}

	return
}

// GetTables 获取课程列表
func (u *AIScheduleUser) GetTable(id int64) (courses []AppCourse, err error) {

	var result GetCourseRet

	_, err = u.Client.R().
		SetQueryParams(map[string]string{
			"userId":   strconv.FormatInt(u.UserID, 10),
			"deviceId": u.DeviceID,
			"ctId":     strconv.FormatInt(id, 10),
		}).
		SetResult(&result).
		Get(AIScheduleTableURL)

	if err != nil {
		return
	}

	if result.Code != 0 {
		err = errors.New(result.Desc)
		return
	}

	courses = result.Data.Courses

	return
}

// DeleteTable 删除课表
func (u *AIScheduleUser) DeleteTable(id int64) (err error) {

	type deleteReq struct {
		SID      int64  `json:"sId"`
		DeviceID string `json:"deviceId"`
		CtID     int64  `json:"ctId"`
		UserID   int64  `json:"userId"`
	}

	req := deleteReq{SID: id, CtID: id, DeviceID: u.DeviceID, UserID: u.UserID}

	var result ChangeRet

	_, err = u.Client.R().
		SetBody(req).
		SetResult(&result).
		Delete(AIScheduleTableURL)

	if err != nil {
		return
	}

	if result.Code != 0 {
		err = errors.New(result.Desc)
	}

	return

}

// CreateTable 创建课表
func (u *AIScheduleUser) CreateTable(name string, isCurrent bool) (id int64, err error) {

	var result CreateRet

	type createReq struct {
		Current  int64  `json:"current"`
		DeviceID string `json:"deviceId"`
		Name     string `json:"name"`
		UserID   int64  `json:"userId"`
	}

	req := createReq{Current: 0, DeviceID: u.DeviceID, Name: name, UserID: u.UserID}

	if isCurrent {
		req.Current = 1
	}

	_, err = u.Client.R().
		SetBody(req).
		SetResult(&result).
		Post(AIScheduleTableURL)

	if err != nil {
		return
	}

	if result.Code != 0 {
		err = errors.New(result.Desc)
		return
	}

	id = result.Data

	return
}

// EditTable 修改课表
func (u *AIScheduleUser) EditTable(name string, id int64, setting TableSetting) (err error) {

	type editReq struct {
		CtID     int64        `json:"ctId"`
		DeviceID string       `json:"deviceId"`
		Name     string       `json:"name"`
		Setting  TableSetting `json:"setting"`
		UserID   int64        `json:"userId"`
	}

	req := editReq{Setting: setting, DeviceID: u.DeviceID, Name: name, UserID: u.UserID, CtID: id}
	req.Setting.ID = id

	var ret ChangeRet

	_, err = u.Client.R().
		SetBody(&req).
		SetResult(&ret).
		Put(AIScheduleTableURL)

	if err != nil {
		return
	}

	if !ret.Data {
		err = errors.New(ret.Desc)
	}

	return
}

// AddCourses 课表添加课程
func (u *AIScheduleUser) AddCourses(id int64, courses []AppCourse) (err error) {

	var ret AddRet

	type addCoursesReq struct {
		Courses  []AppCourse `json:"courses"`
		CtID     int64       `json:"ctId"`
		DeviceID string      `json:"deviceId"`
		UserID   int64       `json:"userId"`
	}

	req := addCoursesReq{Courses: courses, CtID: id, DeviceID: u.DeviceID, UserID: u.UserID}

	_, err = u.Client.R().
		SetBody(&req).
		SetResult(&ret).
		Post(AIScheduleCourseURL)

	if err != nil {
		return
	}

	if ret.Code != 0 {
		err = errors.New(ret.Desc)
		return
	}

	return

}

// ParseTable 解析课表
func (u *AIScheduleUser) ParseTable(tbID int64, html string) (schedule Schedule, err error) {

	type parseReq struct {
		Html string `json:"html"`
		TBID string `json:"tb_id"`
	}

	req := parseReq{Html: strconv.Quote(html), TBID: strconv.FormatInt(tbID, 10)}

	var ret ParseRet

	_, err = u.Client.R().
		SetBody(&req).
		SetResult(&ret).
		Post(AISchedulParserURL)

	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(ret.ParserRet), &schedule)
	if err != nil {
		return
	}

	return schedule, nil

}

// ShareTable 分享课表
func (u *AIScheduleUser) ShareTable(ctId int64) (url string, duetime int64, err error) {

	var ret ShareRet

	_, err = u.Client.R().
		SetQueryParams(map[string]string{
			"ctId":     strconv.FormatInt(ctId, 10),
			"userId":   strconv.FormatInt(u.UserID, 10),
			"deviceId": u.DeviceID,
		}).
		SetResult(&ret).
		Get(AIScheduleShareURL)

	if err != nil {
		return
	}

	if ret.Code != 0 {
		err = errors.New(ret.Desc)
		return
	}

	rawStr := fmt.Sprintf("%d&%s&%s&%d&%d", u.UserID, u.DeviceID, ret.Data.Token, ret.Data.DueTime, ctId)

	url = fmt.Sprintf("%s/#/import_schedule?linkToken=%s", AIScheduleImportURL, base64.StdEncoding.EncodeToString([]byte(rawStr)))

	duetime = ret.Data.DueTime

	return

}

// FeedbackTable 分享课表导入功能 theType {1: "完美", 2: "部分错误", 3: "不可用"}
func (u *AIScheduleUser) FeedbackTable(feedbackType int64, tbID int64) (err error) {

	type feedbackReq struct {
		DeviceID string `json:"deviceId"`
		TbID     int64  `json:"tb_id"`
		Type     int64  `json:"type"`
		UserID   int64  `json:"userId"`
	}

	if feedbackType != 1 && feedbackType != 2 && feedbackType != 3 {
		err = errors.New("feedback type error")
		return
	}

	req := feedbackReq{DeviceID: u.DeviceID, UserID: u.UserID, Type: feedbackType, TbID: tbID}

	_, err = u.Client.R().
		SetBody(&req).
		Post(AIScheduleFeedbackURL)

	return

}
