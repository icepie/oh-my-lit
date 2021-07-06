package sec

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetCurrentMember 获取当前用户信息
func (u *SecUser) GetCurrentMember() (rte CurrentMemberRte, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", GetCurrentMemberUrl+"?vpn-0", nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyText, &rte)
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

// GetStudent 通过学号获取学生信息
func (u *SecUser) GetStudent(stuId string) (rte GetStudentRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("studentId=" + stuId)
	req, err := http.NewRequest("POST", GetStuUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
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

// GetClassmates 通过学号获取学生同班同学信息
func (u *SecUser) GetClassmatesDetail(stuId string) (rte GetClassmatesDetailRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("userName=" + stuId)
	req, err := http.NewRequest("POST", GetClassmatesDetailUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 门户未登陆的情况
	if rte.Count == 0 {
		err = errors.New("no result")
	}

	return
}

// GetClassmates 通过学号获取学生同班同学列表
func (u *SecUser) GetClassmates(stuId string) (rte GetClassmatesRte, err error) {

	client := &http.Client{}

	// 反正一个班没那么多人, 给个理想的值..一次拿完
	pageNum := "1"
	pageSize := "99"

	var data = strings.NewReader("userName=" + stuId + "&pageNum=" + pageNum + "&pageSize=" + pageSize)
	req, err := http.NewRequest("POST", GetClassmatesUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	// 门户未登陆的情况
	if len(rte.Obj) == 0 {
		err = errors.New("no result")
	}

	return
}

// GetOneCardBalance 通过学号获取一卡通剩余金额
func (u *SecUser) GetOneCardBalance(stuId string) (rte GetOneCardBalanceRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("username=" + stuId)
	req, err := http.NewRequest("POST", GetOneCardBalanceUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetOneCardChargeRecords 通过学号获取一卡通充值记录
func (u *SecUser) GetOneCardChargeRecords(stuId string, pageNum uint, pageSize uint) (rte GetOneCardChargeRecordsRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("generaCardRechargeRecordNumber=" + stuId + "&pageNum=" + fmt.Sprint(pageNum) + "&pageSize=" + fmt.Sprint(pageSize))
	req, err := http.NewRequest("POST", GetOneCardChargeRecordsUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetOneCardConsumeRecords 通过学号获取一卡通充值记录
func (u *SecUser) GetOneCardConsumeRecords(stuId string, pageNum uint, pageSize uint) (rte GetOneCardConsumeRecordsRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("generaCardConsumeRecordNumber=" + stuId + "&pageNum=" + fmt.Sprint(pageNum) + "&pageSize=" + fmt.Sprint(pageSize))
	req, err := http.NewRequest("POST", GetOneCardConsumeRecordsUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetExamArrangements 通过学号获取考试安排
func (u *SecUser) GetExamArrangements(stuId string, schoolYear uint, term uint) (rte GetExamArrangementsRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("examinationStudentId=" + stuId + "&examinationSchoolYear=" + fmt.Sprint(schoolYear) + "&examinationTerm=" + fmt.Sprint(term))
	req, err := http.NewRequest("POST", GetExamArrangementsUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}

// GetWeekCourses 通过学号获取获取周课表 role: 学生为 1
func (u *SecUser) GetWeekCourses(stuId string, currentTime string, role uint) (rte GetWeekCoursesRte, err error) {

	client := &http.Client{}

	var data = strings.NewReader("userName=" + stuId + "&currentTime=" + currentTime + "&role=" + fmt.Sprint(role))
	req, err := http.NewRequest("POST", GetWeekCoursesUrl+"?vpn-0", data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", PortalUserUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if !rte.Success {
		err = errors.New(rte.Msg)
	}

	return
}
