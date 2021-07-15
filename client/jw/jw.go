package jw

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/icepie/oh-my-lit/client/sec"
)

var (
	// MianHost
	MianHost = "jw.sec.lit.edu.cn"
	// MianUrl 主网址
	MianUrl = "http://" + MianHost // 120.194.42.205:9001
	// DefaultPath 首页
	DefaultPath = "/default.aspx"
	// LoginPath 登陆地址
	LoginPath = "/_data/index_LOGIN.aspx"
	// MenuPath 菜单地址
	MenuPath = "/frame/menu.aspx"
	// SYSBannerPath 管理员菜单地址
	SYSBannerPath = "/SYS/Main_banner.aspx"
	// MAINFRMPath 主页
	MAINFRMPath = "/MAINFRM.aspx"
	// UserAgent UA
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
	// SchoolCode 院校代号
	SchoolCode = "11070"

	/* 身份
	<select style="WIDTH: 136px" onchange="SelType(this)" name="Sel_Type" id="Sel_Type">
		<option value="STU" usrid="学　号">学生</option>
		<option value="TEA" usrid="工　号">教师教辅人员</option>
		<option value="SYS" usrid="帐　号">管理人员</option>
		<option value="ADM" usrid="帐　号">门户维护员</option>
	</select>
	*/

	// STUType 学生
	STUType = "STU"
	// TEAType 教师教辅人员
	TEAType = "TEA"
	// SYSType 管理人员
	SYSType = "SYS"
	// ADMType 门户维护员
	ADMType = "ADM"
)

// JwUser 教务在线结构体
type JwUser struct {
	Username   string
	Password   string
	SelType    string
	Url        *url.URL
	IsBoundSec bool
	Cookies    []*http.Cookie
}

// NewJwUser 新建教务用户
func NewJwUser(username string, password string, selType string) (user JwUser, err error) {

	if selType != STUType && selType != TEAType && selType != SYSType && selType != ADMType {
		err = errors.New("SelType Error")
		return
	}

	user.Username = username
	user.Password = password
	user.SelType = selType
	user.Url, _ = url.Parse(MianUrl)

	return
}

// JwUser reset main URL
func (u *JwUser) SetUrl(mianUrl string) (err error) {

	mUrl, err := url.Parse(mianUrl)
	if err != nil {
		return
	}

	u.Url = mUrl

	return

}

// 绑定智慧门户帐号
func (u *JwUser) BindSecUser(secUser sec.SecUser) (err error) {

	portalIsLogged := secUser.IsPortalLogged()
	if !portalIsLogged {
		err = errors.New("secUser is not portal logged")
		return
	}

	// https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwOC4xNzMuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5/cas_njjz.aspx?vpn-0

	// 更新Url
	err = u.SetUrl(sec.JWUrlPerfix)
	if err != nil {
		return
	}

	// 覆盖Cookies
	u.Cookies = secUser.Cookies
	u.IsBoundSec = true

	return

}
