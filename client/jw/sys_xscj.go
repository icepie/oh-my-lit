package jw

import (
	"errors"
	"fmt"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetListXSRpt 查询学生
func (u *JwUser) GetListXSRpt(stuID string) (body string, err error) {

	theUrl := u.Url.String() + SysListXSPath

	resp, _ := u.Client.R().
		SetQueryParam("id", stuID).
		SetHeader("referer", theUrl).
		Get(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// GetCXGRCJRpt 获取学生成绩
// timeMode {0: 入学以来, 1: 按照学年, 2: 按照学期}
// mode {0: 原始成绩, 1: 有效成绩}
// flag {0: 主修, 1: 辅修}
func (u *JwUser) GetCXGRCJRpt(timeMode uint, mode uint, flag uint, year uint, term uint, stuID string) (body string, err error) {

	StuMyInfoUrl := u.Url.String() + SysXSGRCJPath

	resp, _ := u.Client.R().
		SetFormData(map[string]string{
			"txt_xm":   u.Username,
			"SelXNXQ":  fmt.Sprint(timeMode),
			"SJ":       fmt.Sprint(mode),
			"sel_xn":   fmt.Sprint(year),
			"sel_xq":   fmt.Sprint(term),
			"zfx_flag": fmt.Sprint(flag),
			"sel_xs":   stuID,
		}).
		SetHeader("referer", StuMyInfoUrl).
		Post(StuMyInfoUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// GetXscjbmLeftRpt 有效成绩排名 left 结果
// timeMode {0: 入学以来, 1: 按照学年, 2: 按照学期}
func (u *JwUser) GetXscjbmLeftRpt(classID string, year uint, term uint, timeMode uint, isBX bool, isXX bool, isRX bool, isHJ bool) (body string, err error) {

	theUrl := u.Url.String() + SysXscjbmLeftPath

	p := map[string]string{
		"sel_bj":  classID,
		"sel_xn":  fmt.Sprint(year),
		"sel_xq":  fmt.Sprint(term),
		"xnxq":    fmt.Sprintf("%d%d", year, term),
		"SelXNXQ": fmt.Sprint(timeMode),
	}

	if isBX {
		p["bxkc"] = "01"
	}

	if isXX {
		p["xxkc"] = "02"
	}

	if isRX {
		p["rxkc"] = "03"
	}

	if isHJ {
		p["hj"] = "1"
	}

	resp, _ := u.Client.R().
		SetFormData(p).
		SetHeader("referer", theUrl).
		Post(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// GetXscjbmRightRpt 有效成绩排名 right 最终结果
// timeMode {0: 入学以来, 1: 按照学年, 2: 按照学期}
func (u *JwUser) GetXscjbmRightRpt(classID string, year uint, term uint, timeMode uint, HidKC string, isBX bool, isXX bool, isRX bool, isHJ bool) (body string, err error) {

	theUrl := u.Url.String() + SysXscjbmRightPath

	p := map[string]string{
		"sel_bj":  classID,
		"sel_xn":  fmt.Sprint(year),
		"sel_xq":  fmt.Sprint(term),
		"hid_kc":  HidKC,
		"xnxq":    fmt.Sprintf("%d%d", year, term),
		"SelXNXQ": fmt.Sprint(timeMode),
	}

	if isBX {
		p["bxkc"] = "01"
	}

	if isXX {
		p["xxkc"] = "02"
	}

	if isRX {
		p["rxkc"] = "03"
	}

	if isHJ {
		p["hj"] = "1"
	}

	resp, _ := u.Client.R().
		SetFormData(p).
		SetHeader("referer", theUrl).
		Post(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// GetXscjbmRpt 有效成绩排名 最终结果
// timeMode {0: 入学以来, 1: 按照学年, 2: 按照学期}
func (u *JwUser) GetXscjbmRpt(classID string, year uint, term uint, timeMode uint, isBX bool, isXX bool, isRX bool, isHJ bool) (body string, err error) {

	b, err := u.GetXscjbmLeftRpt(classID, year, term, timeMode, isBX, isXX, isRX, isHJ)
	if err != nil {
		return
	}

	body, err = u.GetXscjbmRightRpt(classID, year, term, timeMode, u.XscjbmLeftRptToHidKC(b), isBX, isXX, isRX, isHJ)
	if err != nil {
		return
	}

	return
}
