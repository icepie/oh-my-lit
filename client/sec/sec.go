package sec

import (
	"io/ioutil"
	"net/http"

	"github.com/icepie/oh-my-lit/client/util"
)

// 一些常量的定义
var (
	AuthorityUrl = "sec.lit.edu.cn"
	// SecUrl 智慧门户主页
	SecUrl = "https://" + AuthorityUrl
	// AuthPath 认证界面的特殊路径
	AuthPath = "LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy"
	// PortalPath 门户界面的特殊路径
	PortalPath = "LjIwNi4xNzAuMjE4LjE2Mi4xNjg=/LjIxMS4xNzUuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5"
	// AuthlUrlPerfix 认证页面前戳
	AuthlUrlPerfix = SecUrl + "/webvpn/" + AuthPath
	// PortalUrlPerfix 门户页面前戳
	PortalUrlPerfix = SecUrl + "/webvpn/" + PortalPath
	// NeedCaptchaUrl 检查是否需要验证码登陆的接口
	NeedCaptchaUrl = AuthlUrlPerfix + "/authserver/needCaptcha.html"
	// CaptchaUrl 获取验证码
	CaptchaUrl = AuthlUrlPerfix + "/authserver/captcha.html"
	// HomeIndexUrl 导航主页
	HomeIndexUrl = SecUrl + "/frontend_static/frontend/login/index.html"
	// GetHomeParamUrl 主页参数
	GetHomeParamUrl = SecUrl + "/rump_frontend/getHomeParam/"
	// PortalIndexUrl 门户首页
	PortalIndexUrl = PortalUrlPerfix + "/pc/lit/index.html"
	// PortalLoginUrl 门户登陆地址 (第二层)
	PortalLoginUrl = PortalUrlPerfix + "/portal/login/pcLogin"
	// PortalUserUrl 门户个人信息主页
	PortalUserUrl = PortalUrlPerfix + "/portal/pc/lit/user.html"
	// GetCurrentMemberUrl 获取当前门户用户的接口
	GetCurrentMemberUrl = PortalUrlPerfix + "/portal/myCenter/getMemberInfoForCurrentMember"
	// GetStuUrl 获取学生信息接口
	GetStuUrl = PortalUrlPerfix + "/microapplication/api/index/getStudentByStudentId"
	//  GetClassmatesDetail 获取学生同班同学信息接口
	GetClassmatesDetailUrl = PortalUrlPerfix + "/microapplication/api/myclass/findMyclassmatesDetailCount"
	// GetClassmatesUrl 获取学生同班同学列表接口
	GetClassmatesUrl = PortalUrlPerfix + "/microapplication/api/myclass/findMyclassmates"
	// GetOneCardBalanceUrl 获取一卡通余额
	GetOneCardBalanceUrl = PortalUrlPerfix + "/microapplication/api/index/getBalanceAndConsumeThisMonthAndLastMonth"
	// GetOneCardChargeRecordsUrl 获取一卡通充值记录
	GetOneCardChargeRecordsUrl = PortalUrlPerfix + "/microapplication/api/index/listGeneraCardRechargeRecordByGeneraCardRechargeRecordNumberPage"
	// GetOneCardChargeRecordsUrl 获取一卡通消费记录
	GetOneCardConsumeRecordsUrl = PortalUrlPerfix + "/microapplication/api/index/ListGeneraCardConsumeRecordByGeneraCardConsumeRecordNumberPage"
	// // GetDepartmentPhoneList 获取部门电话列表
	// GetDepartmentPhoneList = PortalUrlPerfix + "/microapplication/api/myclass/findMyclassmates"
	// UA
	UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
)

// SecUser 智慧门户用户结构体
type SecUser struct {
	Username string
	Password string
	AuthUrl  string // 真实认证地址 (SecUrl" + "/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F)
	Cookies  []*http.Cookie
}

// NewSecUser 新建智慧门户用户
func NewSecUser(username string, password string) (user SecUser, err error) {

	user.Username = username
	user.Password = password

	// 先从主页拿到真实的登陆地址以及初始化cookies
	client := &http.Client{}

	req, err := http.NewRequest("GET", SecUrl, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	user.AuthUrl, err = util.GetSubstringBetweenStringsByRE(string(bodyText), `<a href="`, `"`)
	if err != nil {
		return
	}

	user.Cookies = resp.Cookies()

	return
}
