package jw

import (
	"errors"
	"fmt"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetClassSelRpt 获取班级课表
func (u *JwUser) GetClassSelRpt(year int64, term int64, calssID string) (body string, err error) {

	StuMyInfoUrl := u.Url.String() + SysClassSelPath

	resp, _ := u.Client.R().
		SetFormData(map[string]string{
			"Sel_XNXQ": fmt.Sprintf("%d%d", year, term),
			"Sel_BJ":   calssID,
		}).
		SetHeader("referer", StuMyInfoUrl).
		Post(StuMyInfoUrl)

	body = util.GB18030ToUTF8(resp.String())

	if strings.Contains(body, "您无权访问此页") {
		err = errors.New("your account does not have permission")
	}

	return
}
