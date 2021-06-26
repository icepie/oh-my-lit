package zhyd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/icepie/oh-my-lit/client/util"
)

// IsNeedCaptcha 判断是否需要验证码登陆
func (u *ZhydUser) IsNeedCaptcha() (isNeed bool, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", NeedCaptchaUrl+"?username="+u.Username+"&_="+fmt.Sprint(time.Now().Unix()), nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("User-Agent", UA)

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
func (u *ZhydUser) GetCaptche() (pix []byte, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", CaptchaUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("User-Agent", UA)

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

// 登陆
func (u *ZhydUser) login(captcha string) (err error) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// 验证码参数
	var captchaParam string

	if len(captcha) > 0 {
		captchaParam = "&captchaResponse=" + captcha
	}

	req, err := http.NewRequest("GET", LoginUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UA)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body := string(bodyText)

	lt, err := util.GetSubstingBetweenStrings(body, `name="lt" value="`, `"`)
	if err != nil {
		return
	}

	execution, err := util.GetSubstingBetweenStrings(body, `name="execution" value="`, `"`)
	if err != nil {
		return
	}

	eventId, err := util.GetSubstingBetweenStrings(body, `name="_eventId" value="`, `"`)
	if err != nil {
		return
	}

	rmShown, err := util.GetSubstingBetweenStrings(body, ` name="rmShown" value="`, `"`)
	if err != nil {
		return
	}

	//log.Println(lt, execution, eventId, rmShown)

	// u.Cookies = resp.Cookies()

	// 开始登陆
	var data = strings.NewReader("username=" + u.Username + "&password=" + u.Password + captchaParam + "&lt=" + lt + "&execution=" + execution + "&_eventId=" + eventId + "&rmShown=" + rmShown)

	req, err = http.NewRequest("POST", LoginUrl, data)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Content-Length", "133")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", UA)

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(bodyText)

	log.Println(body)

	location, err := resp.Location()
	if err != nil {
		return
	}

	if location.String() != "http://ids.lit.edu.cn/authserver/userAttributesEdit.do" {
		err = errors.New("login error")
		return
	}

	// 添加这玩意
	u.Cookies = append(u.Cookies, resp.Cookies()...)

	// 判断是否有错误
	// if strings.Contains(body, "callback_err_login") {
	// 	loginErrStr, _ := util.GetSubstingBetweenStrings(body, `callback_err_login">`, `</div>`)

	// 	err = errors.New(loginErrStr)

	// 	return
	// } else if strings.Contains(body, "login_hint") {
	// 	err = errors.New("login error")

	// 	return
	// }

	return

}

// Login 普通登陆
func (u *ZhydUser) Login() (err error) {
	return u.login("")
}

// LoginWithCap 验证码登陆
func (u *ZhydUser) LoginWithCap(captcha string) (err error) {
	return u.login(captcha)
}
