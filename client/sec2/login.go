package sec2

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/icepie/oh-my-lit/client/util"
)

// IsLogged() 检测用户是否登陆
func (u *SecUser) IsLogged() (isLogged bool) {

	_, err := u.GetHomeParam()

	return err == nil
}

// IsNeedCaptcha 判断是否需要验证码登陆
func (u *SecUser) IsNeedCaptcha() (isNeed bool, err error) {

	resp, reqErr := u.Client.R().
		SetQueryParams(map[string]string{
			"username": u.Username,
			"_":        fmt.Sprint(time.Now().Unix()),
		}).
		SetHeader("referer", u.AuthUrl).
		Get(u.AuthlUrlPerfix + NeedCaptchaPath)

	if resp.StatusCode() != 200 {
		err = reqErr
		return
	}

	body := resp.String()

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

	resp, err := u.Client.R().
		SetQueryParams(map[string]string{
			"username": u.Username,
			"_":        fmt.Sprint(time.Now().Unix()),
		}).
		SetHeader("referer", u.AuthUrl).
		SetHeader("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8").
		Get(u.AuthlUrlPerfix + CaptchaPath)

	if err != nil {
		return
	}

	pix = resp.Body()

	return

}

// login 通用登陆
func (u *SecUser) login(captcha string) (err error) {

	if len(u.Username) == 0 {
		err = errors.New("empty username")
		return
	}

	if len(u.Password) == 0 {
		err = errors.New("empty password")
		return
	}

	// 刷新 webvpn path
	err = u.PerSetCooikes()
	if err != nil {
		return
	}

	// client := &http.Client{}

	resp, _ := u.Client.R().
		SetHeader("referer", u.AuthUrl).
		Get(u.AuthUrl)

	body := resp.String()

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

	// var data = strings.NewReader("username=" + u.Username + "&password=" + u.Password + captchaParam + "&lt=" + lt + "&execution=" + execution + "&_eventId=" + eventId + "&rmShown=" + rmShown)

	req := u.Client.R().
		SetHeader("referer", u.AuthUrl).
		SetHeader("authority", actionUrl).
		SetFormData(map[string]string{
			"username":  u.Username,
			"password":  u.Password,
			"lt":        lt,
			"execution": execution,
			"_eventId":  eventId,
			"rmShown":   rmShown,
		})

	if len(captcha) > 0 {
		req.SetFormData(map[string]string{
			"captchaResponse": captcha,
		})
	}

	resp, _ = req.Post(decodeurl)
	if err != nil {
		return
	}

	body = resp.String()

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

	// // 获取门户path
	// u.getPortalPath()

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
