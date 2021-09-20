package jw

import (
	"errors"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/icepie/oh-my-lit/client/sec"
)

var (
	// MianHost
	MianHost = "jw.sec.lit.edu.cn"
	// MianUrl 主网址
	MianUrl = "http://" + MianHost // 120.194.42.205:9001
	// DefaultPath 首页
	DefaultPath = "/default.aspx"
	// LoginPath 登陆接口
	LoginPath = "/_data/index_LOGIN.aspx"
	// MenuPath 菜单接口
	MenuPath = "/frame/menu.aspx"
	// SysBannerPath 管理员菜单接口(通用)
	SysBannerPath = "/SYS/Main_banner.aspx"
	// StuMyInfoPath 学生学籍信息接口
	StuMyInfoPath = "/xsxj/Stu_MyInfo_RPT.aspx"
	// StuZXJGPath 学生正选结果
	StuZXJGPath = "/wsxk/stu_zxjg_rpt.aspx"
	// StuBYFAKCPath 学生理论课程结果
	StuBYFAKCPath = "/jxjh/Stu_byfakc_rpt.aspx"
	// StuBYFAHJPath 学生实践环节结果
	StuBYFAHJPath = "/jxjh/Stu_byfahj_rpt.aspx"
	// SysListXSPath 管理员查询学生姓名结果
	SysListXSPath = "/XSCJ/Private/list_XS.aspx"
	// SysXSGRCJPath 管理员查询学生成绩结果
	SysXSGRCJPath = "/XSCJ/f_xsgrcj_rpt.aspx"
	// SysClassSelPath 管理员查询班级课表
	SysClassSelPath = "/ZNPK/ClassSel_rpt.aspx"
	// ClassSelPath 公用课表接口
	ClassSelPath = "/ZNPK/KBFB_ClassSel.aspx"
	// LoginBySecPath 智慧门户快速登陆
	LoginBySecPath = "/cas_njjz.aspx"
	// MAINFRMPath 主页
	MAINFRMPath = "/MAINFRM.aspx"
	// UA
	UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	// SchoolCode 院校代号
	SchoolCode = "11070"

	// MainHeaders 主请求头
	MainHeaders = map[string]string{
		"dnt":              "1",
		"x-requested-with": "XMLHttpRequest",
		"sec-ch-ua-mobile": "1",
		"User-Agent":       UA,
		"sec-fetch-site":   "same-origin",
		"sec-fetch-mode":   "cors",
		"sec-fetch-dest":   "empty",
		"Accept-Language":  "zh-CN,zh;q=0.9",
	}

	/* 身份
	<select style="WIDTH: 136px" onchange="SelType(this)" name="Sel_Type" id="Sel_Type">
		<option value="STU" usrid="学　号">学生</option>
		<option value="TEA" usrid="工　号">教师教辅人员</option>
		<option value="SYS" usrid="帐　号">管理人员</option>
		<option value="ADM" usrid="帐　号">门户维护员</option>
	</select>
	*/

	// StuType 学生
	StuType = "STU"
	// TeaType 教师教辅人员
	TeaType = "TEA"
	// SysType 管理人员
	SysType = "SYS"
	// AdmType 门户维护员
	AdmType = "ADM"
)

// JwUser 教务在线结构体
type JwUser struct {
	Username   string
	Password   string
	SelType    string
	Url        *url.URL
	IsBoundSec bool
	Client     *resty.Client
}

// NewJwUser 新建教务用户
func NewJwUser() *JwUser {
	var u JwUser
	u.Client = resty.New()
	u.Client.SetHeaders(MainHeaders)
	u.Url, _ = url.Parse(MianUrl)
	return &u
}

// SetUsername 设置用户名
func (u *JwUser) SetUsername(username string) *JwUser {
	u.Username = username
	return u
}

// SetPassword 设置密码
func (u *JwUser) SetPassword(password string) *JwUser {
	u.Password = password
	return u
}

// SetSelType 设置登陆类型
/*
	// StuType 学生
	StuType = "STU"
	// TeaType 教师教辅人员
	TeaType = "TEA"
	// SysType 管理人员
	SysType = "SYS"
	// AdmType 门户维护员
	AdmType = "ADM"
*/
func (u *JwUser) SetSelType(selType string) *JwUser {
	if selType != StuType && selType != TeaType && selType != SysType && selType != AdmType {
		// log.Println("error: your selType not support!")
		return u
	}
	u.SelType = selType
	return u
}

// SetUrl JwUser reset main URL
func (u *JwUser) SetUrl(url *url.URL) *JwUser {
	u.Url = url
	return u
}

// BindSecUser 绑定智慧门户帐号
func (u *JwUser) BindSecUser(secUser *sec.SecUser) (err error) {

	portalIsLogged := secUser.IsPortalLogged()
	if !portalIsLogged {
		err = errors.New("secUser is not portal logged")
		return
	}

	// https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwOC4xNzMuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5/cas_njjz.aspx?vpn-0

	secJWUrl, _ := url.Parse(sec.JWUrlPerfix)

	// 更新Url
	u.SetUrl(secJWUrl)

	// // 覆盖Client
	u.Client.SetCookies(secUser.Client.GetClient().Jar.Cookies(secJWUrl))
	// u.Client = secUser.Client
	// u.Client.SetDebug(true)
	u.IsBoundSec = true

	return

}
