package sec2

import (
	"encoding/json"
	"errors"
)

// GetCurrentMember 获取当前用户信息
func (u *SecUser) GetCurrentMember() (rte CurrentMemberRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("accept", "application/json, text/plain, */*").
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		Get(u.PortalUrlPerfix + GetCurrentMemberPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 门户未登陆的情况
	if len(rte.Obj.MemberID) == 0 {
		err = errors.New("portal not be logged")
	}

	return
}

// GetStudentByStuID 通过学号获取学生信息
func (u *SecUser) GetStudentByStuID(stuID string) (rte GetStudentRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"studentId": stuID,
		}).
		Post(u.PortalUrlPerfix + GetStuPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 门户未登陆的情况
	if len(rte.Obj.StudentID) == 0 {
		err = errors.New("no result")
	}

	return
}

// GetStudent 获取学生信息
func (u *SecUser) GetStuden() (GetStudentRte, error) {
	return u.GetStudentByStuID(u.Username)
}
