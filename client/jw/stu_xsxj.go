package jw

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetStuMyInfoRpt 获取学生学籍信息原数据
func (u *JwUser) GetStuMyInfoRpt() (body string, err error) {

	StuMyInfoUrl := u.Url.String() + StuMyInfoPath

	resp, _ := u.Client.R().
		SetHeader("referer", StuMyInfoUrl).
		Get(StuMyInfoUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
