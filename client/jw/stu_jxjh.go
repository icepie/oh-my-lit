package jw

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetStuXKJGRpt 获取学生理论课程原始数据
func (u *JwUser) GetStuBYFAKCRpt() (body string, err error) {

	theUrl := u.Url.String() + StuBYFAKCPath

	resp, _ := u.Client.R().
		SetHeader("referer", theUrl).
		Get(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}

// GetStuXKJGRpt 获取学生教学方案原始数据
func (u *JwUser) GetStuBYFAHJRpt() (body string, err error) {

	StuBYFAHJUrl := u.Url.String() + StuBYFAHJPath

	resp, _ := u.Client.R().
		SetHeader("referer", StuBYFAHJUrl).
		Get(StuBYFAHJUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
