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
	SecUrl = "https://" + AuthorityUrl + "/"
	// AuthPath 认证界面的特殊路径
	AuthPath = "LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy"
	// NeedCaptchaUrl 检查是否需要验证码登陆的接口
	NeedCaptchaUrl = "https://sec.lit.edu.cn/webvpn/" + AuthPath + "/authserver/needCaptcha.html"
	// CaptchaUrl 获取验证码
	CaptchaUrl = "https://sec.lit.edu.cn/webvpn/" + AuthPath + "/authserver/captcha.html"
	// UA
	UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
)

// SecUser 智慧门户用户结构体
type SecUser struct {
	Username   string
	Password   string
	IsPcLogged bool   // 是否进行第二层登陆
	IsLogged   bool   // 是否登陆
	AuthUrl    string // 真实认证地址 (https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F)
	Cookies    []*http.Cookie
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

	user.AuthUrl, err = util.GetSubstingBetweenStrings(string(bodyText), `<a href="`, `"`)
	if err != nil {
		return
	}

	user.Cookies = resp.Cookies()

	return
}
