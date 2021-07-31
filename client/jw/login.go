package jw

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/icepie/oh-my-lit/client/util"
)

// JwUser Login
func (u *JwUser) Login() (err error) {

	if len(u.Username) == 0 {
		err = errors.New("empty username")
		return
	}

	if len(u.Password) == 0 {
		err = errors.New("empty password")
		return
	}

	if len(u.SelType) == 0 {
		err = errors.New("empty selType")
		return
	}

	// if u.IsBoundSec {
	// 	// 暂时未实现登陆任意帐号
	// 	err = errors.New("please use LoginBySec()")
	// 	return
	// }

	LoginUrl := u.Url.String() + LoginPath

	resp, _ := u.Client.R().
		SetHeader("referer", LoginUrl).
		Get(LoginUrl)

	// 将 gb2312 转换为 utf-8
	body := util.GB18030ToUTF8(resp.String())

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return
	}

	// 拿到参数
	vs, isExist := doc.Find("input[name=__VIEWSTATE]").First().Attr("value")
	if !isExist {
		err = errors.New("vs is no exist")
		return
	}

	// the last step
	u.Client.R().
		SetFormData(map[string]string{
			"__VIEWSTATE":             vs,
			"Sel_Type":                u.SelType, // SYS etc..
			"txt_sdsdfdsfryuiighgdf":  u.Username,
			"txt_dsfdgtjhjuixssdsdf":  "",
			"txt_sftfgtrefjdndcfgerg": "",
			"typeName":                "",
			"sdfdfdhgwerewt":          chkpwd(u.Username, u.Password),
			"cxfdsfdshjhjlk":          "",
		}).
		SetHeader("referer", LoginUrl).
		Post(LoginUrl)

	isLogged := u.IsLogged()

	// log.Println(resp.String())

	if !isLogged {
		err = errors.New("jw fail to login")
	}

	return

}

// LoginBySec 通过智慧门户帐号快速登陆
func (u *JwUser) LoginBySec() (err error) {

	if !u.IsBoundSec {
		// 暂时未实现登陆任意帐号
		err = errors.New("only for sec user")
		return
	}

	if len(u.SelType) == 0 {
		err = errors.New("empty selType")
		return
	}

	LoginBySecUrl := u.Url.String() + LoginBySecPath + "?vpn-0"

	resp, _ := u.Client.R().
		SetHeader("referer", LoginBySecUrl).
		Get(LoginBySecUrl)

	body := util.GB18030ToUTF8(resp.String())

	isMultiRole := false
	if strings.Contains(body, "请选择相应角色") {
		isMultiRole = true
	}

	isHaveRole := false
	if isMultiRole {
		sel_Type := ""
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
		doc.Find("select#Sel_Grp>option").Each(func(i int, option *goquery.Selection) {
			op, _ := option.Attr("value")
			if op[:3] == u.SelType {
				isHaveRole = true
				sel_Type = op
			}
		})

		userID, _ := doc.Find("input#UserID").First().Attr("value")
		roleCHK, _ := doc.Find("input#roleCHK").First().Attr("value")
		typeName, _ := doc.Find("input#typeName").First().Attr("value")
		pcInfo, _ := doc.Find("input#pcInfo").First().Attr("value")

		// the last step
		u.Client.R().
			SetFormData(map[string]string{
				"Sel_Type": sel_Type,
				"subBtn1":  "%C8%B7%B6%A8",
				"vUserID":  userID,
				"roleCHK":  roleCHK,
				"typeName": typeName,
				"pcInfo":   pcInfo,
			}).
			SetHeader("referer", LoginBySecUrl).
			Post(LoginBySecUrl)

	}

	if isMultiRole && !isHaveRole {
		err = errors.New("you do not have the role by the selType")
		return
	}

	isLogged := u.IsLogged()

	if !isLogged {
		err = errors.New("jw fail to login")
	}

	return

}

// 是否登陆
func (u *JwUser) IsLogged() (isLogged bool) {

	isLogged = false

	MAINFRMUrl := u.Url.String() + MenuPath

	if u.IsBoundSec {
		MAINFRMUrl += "?vpn-0"
	}

	resp, _ := u.Client.R().
		SetHeader("referer", MAINFRMUrl).
		Get(MAINFRMUrl)

	body := util.GB18030ToUTF8(resp.String())

	// 检测是否登陆成功
	if strings.Contains(body, "洛阳理工学院教务") {
		isLogged = true
	}

	return
}
