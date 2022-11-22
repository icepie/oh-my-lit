package jw

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetStuDJKSCJ 获取学生等级成绩
func (u *JwUser) GetStuDJKSCJ() (body string, err error) {

	StuZXJGUrl := u.Url.String() + StuDJKSCJPath

	resp, _ := u.Client.R().
		SetHeader("referer", StuZXJGUrl).
		Get(StuZXJGUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
