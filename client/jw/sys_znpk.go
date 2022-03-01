package jw

import (
	"errors"
	"fmt"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

var (
	// 王城
	WangchengCampus uint = 1
	// 开元
	KaiyuanCampus uint = 2
	// 九都
	JiuduCampus uint = 3
)

type DayfreeSelParam struct {
	Year     uint
	Term     uint
	Week     uint
	Day      uint
	Section  uint
	Campus   uint
	Building uint // 教学楼
	RoomType uint // 教室类型
}

// GetClassSelRpt 获取班级课表
func (u *JwUser) GetClassSelRpt(year uint, term uint, calssID string) (body string, err error) {

	theUrl := u.Url.String() + SysClassSelPath

	resp, _ := u.Client.R().
		SetFormData(map[string]string{
			"Sel_XNXQ": fmt.Sprintf("%d%d", year, term),
			"Sel_BJ":   calssID,
		}).
		SetHeader("referer", theUrl).
		Post(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// 获取空教室
func (u *JwUser) GetDayfreeSelRpt(data DayfreeSelParam) (body string, err error) {

	theUrl := u.Url.String() + SysZNPKDayfreeSelPath

	formData := map[string]string{
		"SelXN":        fmt.Sprintf("%d%d", data.Year, data.Term),
		"Sel_ZC":       buildNumParam(data.Week),    // 周次
		"selxqs":       fmt.Sprintf("%d", data.Day), // 星期
		"Sel_JC":       buildNumParam(data.Section), // 节次
		"Sel_JXL":      buildNumParam(data.Building),
		"sel_jslx":     buildNumParam(data.RoomType),
		"Submit01":     "%BC%EC%CB%F7",
		"sel_roomname": "",
	}

	if data.Campus != 0 {
		formData["sel_xq"] = buildNumParam(data.Campus)
	}

	resp, _ := u.Client.R().
		SetFormData(formData).
		SetHeader("referer", theUrl).
		Post(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

func buildNumParam(num uint) string {
	if num == 0 {
		return ""
	}
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}
