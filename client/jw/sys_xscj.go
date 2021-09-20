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
