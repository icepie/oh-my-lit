package sec

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

	var data1 = strings.NewReader("studentId=" + stuId)
	req, err := http.NewRequest("POST", "https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mi4xNjg=/LjIxMS4xNzUuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5/microapplication/api/index/getStudentByStudentId?vpn-0", data1)
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
		log.Fatal(err)
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
