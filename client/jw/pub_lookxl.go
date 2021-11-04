package jw

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetLookxlRpt 获取校历返回页面
func (u *JwUser) GetLookxlRpt() (body string, err error) {

	theUrl := u.Url.String() + LookxlRptPath

	resp, _ := u.Client.R().
		SetHeader("referer", theUrl).
		Get(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
