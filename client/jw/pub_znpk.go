package jw

import (
	"github.com/icepie/oh-my-lit/client/util"
)

// GetClassSelRpt 获取公共班级课表页面
func (u *JwUser) GetClassSelPage() (body string, err error) {

	theUrl := u.Url.String() + ClassSelPath

	resp, _ := u.Client.R().
		SetHeader("referer", theUrl).
		Get(theUrl)

	body = util.GB18030ToUTF8(resp.String())

	// if strings.Contains(body, "您无权访问此页") {
	// 	err = errors.New("your account does not have permission")
	// }

	return
}
