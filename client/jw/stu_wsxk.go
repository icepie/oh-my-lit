package jw

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetStuXKJGRpt 获取学生正选结果原数据
func (u *JwUser) GetStuXKJGRpt() (body string, err error) {

	StuZXJGUrl := u.Url.String() + StuZXJGPath

	resp, _ := u.Client.R().
		SetHeader("referer", StuZXJGUrl).
		Get(StuZXJGUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
