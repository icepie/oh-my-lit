package sec

import (
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/icepie/oh-my-lit/client/util"
)

// IsLogged() 检测用户是否登陆
func (u *SecUser) IsLogged() (isLogged bool) {

	_, err := u.GetHomeParam()

	//log.Println(t)

	if err != nil {
		isLogged = false
	} else {
		isLogged = true
		return
	}

	return
}

// IsPortalLogged 是否门户登陆
func (u *SecUser) IsPortalLogged() (isLogged bool) {
	_, err := u.GetCurrentMember()

	if err != nil {
		isLogged = false
	} else {
		isLogged = true
		return
	}

	return

}

// IsNeedCaptcha 判断是否需要验证码登陆
func (u *SecUser) IsNeedCaptcha() (isNeed bool, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.AuthlUrlPerfix+NeedCaptchaPath+"?username="+u.Username+"&_="+fmt.Sprint(time.Now().Unix()), nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	//req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	//req.Header.Set("referer", "https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?vpn-0&amp;amp=&amp;amp;service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F&amp;amp;vpn-0=")
	req.Header.Set("referer", u.AuthUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, _ := ioutil.ReadAll(resp.Body)

	body := string(bodyText)

	// 最后判断是否需要验证码进行登陆
	if strings.HasPrefix(body, "false") {
		isNeed = false
	} else if strings.HasPrefix(body, "true") {
		isNeed = true
	} else {
		err = errors.New("can not get the info")
	}

	return

}

// GetCaptche 获取验证码 (JPEG)
func (u *SecUser) GetCaptche() (pix []byte, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.AuthlUrlPerfix+CaptchaPath, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	//req.Header.Set("sec-ch-ua", " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"")
	req.Header.Set("dnt", "1")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "no-cors")
	req.Header.Set("sec-fetch-dest", "image")
	req.Header.Set("referer", u.AuthUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	//req2.Header.Set("cookie", "csrftoken=jz331d3MpxsWBEFHSW9Scy0v18U6VBzT7NC66xoAvitxV4NkqBpcvF81kytnTe2I; client_vpn_ticket=F8fNTTQFxE8jaUtu; sessionid=ffk4zsykvth2id94ob8g97h6j75bv0ic")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	pix, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return

}

// login 通用登陆
func (u *SecUser) login(captcha string) (err error) {

	// 刷新 webvpn path
	u.prepare()

	client := &http.Client{}

	// 获取必要参数
	req, err := http.NewRequest("GET", u.AuthUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	// 验证码参数
	var captchaParam string

	if len(captcha) > 0 {
		captchaParam = "&captchaResponse=" + captcha
	}

	req.Header.Set("authority", AuthorityUrl)
	//req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("dnt", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	//req.Header.Set("referer", "https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?vpn-0&amp;amp=&amp;amp;service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F&amp;amp;vpn-0=")
	req.Header.Set("referer", u.AuthUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, _ := ioutil.ReadAll(resp.Body)

	body := string(bodyText)

	// 获取所有可需参数
	actionUrl, err := util.GetSubstringBetweenStringsByRE(body, `id="casLoginForm" class="fm-v clearfix" action="`, `"`)
	if err != nil {
		return
	}

	lt, err := util.GetSubstringBetweenStringsByRE(body, `name="lt" value="`, `"`)
	if err != nil {
		return
	}

	execution, err := util.GetSubstringBetweenStringsByRE(body, `name="execution" value="`, `"`)
	if err != nil {
		return
	}

	eventId, err := util.GetSubstringBetweenStringsByRE(body, `name="_eventId" value="`, `"`)
	if err != nil {
		return
	}

	rmShown, err := util.GetSubstringBetweenStringsByRE(body, `name="rmShown" value="`, `"`)
	if err != nil {
		return
	}

	//log.Println(actionUrl, lt, execution, eventId, rmShown)

	// 这个地址需要html解码
	decodeurl := html.UnescapeString(actionUrl)

	var data = strings.NewReader("username=" + u.Username + "&password=" + u.Password + captchaParam + "&lt=" + lt + "&execution=" + execution + "&_eventId=" + eventId + "&rmShown=" + rmShown)

	req, err = http.NewRequest("POST", decodeurl, data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", actionUrl)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("origin", SecUrl)
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("dnt", "1")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("user-agent", UA)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("referer", actionUrl)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, _ = ioutil.ReadAll(resp.Body)

	body = string(bodyText)

	// 判断是否有错误
	if strings.Contains(body, "callback_err_login") {
		loginErrStr, _ := util.GetSubstringBetweenStringsByRE(body, `callback_err_login">`, `</div>`)

		err = errors.New(loginErrStr)

		return
	}

	// 确保账号登陆成功
	if !u.IsLogged() {
		u.login(captcha)
	}

	// 获取门户path
	u.getPortalPath()

	return
}

// Login 第一层普通登陆
func (u *SecUser) Login() (err error) {
	return u.login("")
}

// LoginWithCap 第一层验证码登陆
func (u *SecUser) LoginWithCap(captcha string) (err error) {
	return u.login(captcha)
}

// PortalLogin 第二层门户登陆
func (u *SecUser) PortalLogin() (err error) {

	if len(u.PortalUrlPerfix) == 0 {
		err = errors.New("please login first")
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.PortalUrlPerfix+PortalLoginPath+"?vpn-0", nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("dnt", "1")
	req.Header.Set("user-agent", UA)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("referer", u.PortalUrlPerfix+PortalIndexPath)
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// location, err := resp.Location()
	// if err != nil {
	// 	return
	// }

	// log.Println(location)

	// bodyText, _ := ioutil.ReadAll(resp.Body)

	// body := string(bodyText)

	// log.Println(body)

	// 确保账号登陆成功
	if !u.IsPortalLogged() {
		err = errors.New("fail to login")
	}

	return
}
