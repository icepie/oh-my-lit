package jw

import (
	"errors"
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
