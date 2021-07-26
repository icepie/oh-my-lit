package sec2

import (
	"encoding/json"
	"errors"
	"strconv"
)

// GetCurrentMember 获取当前用户信息
func (u *SecUser) GetCurrentMember() (rte CurrentMemberRte, err error) {

	resp, _ := u.Client.R().
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
func (u *SecUser) GetStudent() (GetStudentRte, error) {
	return u.GetStudentByStuID(u.Username)
}

// GetClassmatesByStuI 通过学号获取学生同班同学信息
func (u *SecUser) GetClassmatesDetailByStuID(stuID string) (rte GetClassmatesDetailRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"userName": stuID,
		}).
		Post(u.PortalUrlPerfix + GetClassmatesDetailPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetClassmatesDetail 获取同班同学详情
func (u *SecUser) GetClassmatesDetail() (GetClassmatesDetailRte, error) {
	return u.GetClassmatesDetailByStuID(u.Username)
}

// GetClassmatesByStuID 通过学号获取学生同班同学列表
//	pageNum := "1"
//	pageSize := "99"
func (u *SecUser) GetClassmatesByStuID(stuID string, pageNum int, pageSize int) (rte GetClassmatesRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"userName": stuID,
			"pageNum":  strconv.Itoa(pageNum),
			"pageSize": strconv.Itoa(pageSize),
		}).
		Post(u.PortalUrlPerfix + GetClassmatesPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetClassmates 获取同班同学
func (u *SecUser) GetClassmates(pageNum int, pageSize int) (GetClassmatesRte, error) {
	return u.GetClassmatesByStuID(u.Username, pageNum, pageSize)
}

// GetWeekCoursesByID 通过账号获取获取周课表
// currentTime: 如 2021-10-29
// role: {学生:1, 教师:2}
func (u *SecUser) GetWeekCoursesByID(username string, currentTime string, role int) (rte GetWeekCoursesRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"userName":    username,
			"currentTime": currentTime,
			"role":        strconv.Itoa(role),
		}).
		Post(u.PortalUrlPerfix + GetWeekCoursesPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetWeekCourses 获取周课表安排
// currentTime: 如 2021-10-29
// role: {学生:1, 教师:2}
func (u *SecUser) GetWeekCourses(currentTime string, role int) (rte GetWeekCoursesRte, err error) {
	return u.GetWeekCoursesByID(u.Username, currentTime, role)
}

// GetExamArrangementsByStuID 通过学号获考试安排
//	schoolYear := "2021" //  学年
//	schoolTerm := "1" // {第一学期:0,第二学期:1}
func (u *SecUser) GetExamArrangementsByStuID(stuID string, schoolYear int, schoolTerm int) (rte GetExamArrangementsRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"examinationStudentId":  stuID,
			"examinationSchoolYear": strconv.Itoa(schoolYear),
			"examinationTerm":       strconv.Itoa(schoolTerm),
		}).
		Post(u.PortalUrlPerfix + GetExamArrangementsPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetExamArrangement 获取考试安排
//	schoolYear := "2021" //  学年
//	schoolTerm := "1" // {第一学期:0,第二学期:1}
func (u *SecUser) GetExamArrangemen(schoolYear int, schoolTerm int) (rte GetExamArrangementsRte, err error) {
	return u.GetExamArrangementsByStuID(u.Username, schoolYear, schoolTerm)
}

// GetOneCardConsumeRecordsByID 通过帐号获取一卡通充值记录
//	pageNum := "1"
//	pageSize := "99"
func (u *SecUser) GetOneCardConsumeRecordsByID(ID string, pageNum int, pageSize int) (rte GetOneCardConsumeRecordsRte, err error) { // GetOneCardConsumeRecords 通过学号获取一卡通充值记录

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"generaCardConsumeRecordNumber": ID,
			"pageNum":                       strconv.Itoa(pageNum),
			"pageSize":                      strconv.Itoa(pageSize),
		}).
		Post(u.PortalUrlPerfix + GetOneCardConsumeRecordsPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetOneCardConsumeRecords 通帐号获取一卡通充值记录
//	pageNum := "1"
//	pageSize := "99"
func (u *SecUser) GetOneCardConsumeRecords(pageNum int, pageSize int) (rte GetOneCardConsumeRecordsRte, err error) { // GetOneCardConsumeRecords 通过学号获取一卡通充值记录
	return u.GetOneCardConsumeRecordsByID(u.Username, pageNum, pageSize)
}

// GetOneCardChargeRecordsByID 通过帐号获取一卡通充值记录
//	pageNum := "1"
//	pageSize := "99"
func (u *SecUser) GetOneCardChargeRecordsByID(ID string, pageNum int, pageSize int) (rte GetOneCardChargeRecordsRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"generaCardRechargeRecordNumber": ID,
			"pageNum":                        strconv.Itoa(pageNum),
			"pageSize":                       strconv.Itoa(pageSize),
		}).
		Post(u.PortalUrlPerfix + GetOneCardChargeRecordsPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return

}

// GetOneCardConsumeRecords 通帐号获取一卡通充值记录
//	pageNum := "1"
//	pageSize := "99"
func (u *SecUser) GetOneCardChargeRecords(pageNum int, pageSize int) (rte GetOneCardChargeRecordsRte, err error) { // GetOneCardConsumeRecords 通过学号获取一卡通充值记录
	return u.GetOneCardChargeRecordsByID(u.Username, pageNum, pageSize)
}

// GetOneCardBalanceByID 通过帐号获取一卡通剩余金额
func (u *SecUser) GetOneCardBalanceByID(ID string) (rte GetOneCardBalanceRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"username": ID,
		}).
		Post(u.PortalUrlPerfix + GetOneCardBalancePath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return

}

// GetOneCardBalance 获取一卡通剩余金额
func (u *SecUser) GetOneCardBalance() (rte GetOneCardBalanceRte, err error) {
	return u.GetOneCardBalanceByID(u.Username)
}

// GetStaffByStaffID 通过职工号获取教职工信息
func (u *SecUser) GetStaffByStaffID(staffID string) (rte GetStaffRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"staffNumber": staffID,
		}).
		Post(u.PortalUrlPerfix + GetStaffPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 门户未登陆的情况
	if len(rte.Obj.StaffName) == 0 {
		err = errors.New("no result")
	}

	return
}

// GetStaff 获取教职工信息
func (u *SecUser) GetStaff() (rte GetStaffRte, err error) {
	return u.GetStaffByStaffID(u.Username)
}

// GetClassStudents 获取班级学生
func (u *SecUser) GetClassStudents(classCode string, pageNum int, pageSize int) (rte GetClassStudentsRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"classcode": classCode,
			"pageNum":   strconv.Itoa(pageNum),
			"pageSize":  strconv.Itoa(pageSize),
		}).
		Post(u.PortalUrlPerfix + GetClassStudentsPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
		return
	}

	// if len(rte.Obj) == 0 {
	// 	err = errors.New("no result")
	// }

	return
}

// GetAllInvigilateByStaffID 通过职工号获取监考安排
//	schoolYear := "2021" //  学年
//	schoolTerm := "1" // {第一学期:0,第二学期:1}
func (u *SecUser) GetAllInvigilateByStaffID(StaffID string, schoolYear int, schoolTerm int) (rte GetAllInvigilateRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"invigilateStaffNumber": StaffID,
			"invigilateSchoolYear":  strconv.Itoa(schoolYear),
			"invigilateTerm":        strconv.Itoa(schoolTerm),
		}).
		Post(u.PortalUrlPerfix + GetAllInvigilatePath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetAllInvigilate 获取监考安排
//	schoolYear := "2021" //  学年
//	schoolTerm := "1" // {第一学期:0,第二学期:1}
func (u *SecUser) GetAllInvigilate(schoolYear int, schoolTerm int) (rte GetAllInvigilateRte, err error) {
	return u.GetAllInvigilateByStaffID(u.Username, schoolYear, schoolTerm)
}

// GetAssetsByStaffID 通过职工号获取资产
func (u *SecUser) GetAssetsByStaffID(staffID string, pageNum int, pageSize int) (rte GetAssetsRte, err error) {

	resp, _ := u.Client.R().
		SetHeader("referer", u.PortalUrlPerfix+PortalUserPath).
		SetFormData(map[string]string{
			"assetsStaffNumber": staffID,
			"pageNum":           strconv.Itoa(pageNum),
			"pageSize":          strconv.Itoa(pageSize),
		}).
		Post(u.PortalUrlPerfix + GetAssetsPath + "?vpn-0")

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetAssetsByStaff 获取资产
func (u *SecUser) GetAssets(pageNum int, pageSize int) (rte GetAssetsRte, err error) {
	return u.GetAssetsByStaffID(u.Username, pageNum, pageSize)
}
