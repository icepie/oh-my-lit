package sec

import (
	"errors"
	"html"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/icepie/oh-my-lit/client/util"
)

// IsLogged() 用户是否已登陆
func (u *SecUser) IsLogged() (isLogged bool) {

	_, err := u.GetHomeParam()

	return err == nil
}

// IsPortalLogged 用户是否已登陆门户
func (u *SecUser) IsPortalLogged() (isLogged bool) {

	_, err := u.GetCurrentMember()

	return err == nil
}

// // IsNeedCaptcha 判断是否需要验证码登陆
// func (u *SecUser) IsNeedCaptcha() (isNeed bool, err error) {

// 	resp, reqErr := u.Client.R().
// 		SetQueryParams(map[string]string{
// 			"username": u.Username,
// 			"_":        fmt.Sprint(time.Now().Unix()),
// 		}).
// 		SetHeader("referer", u.AuthUrl).
// 		Get(u.AuthUrlPerfix + NeedCaptchaPath)

// 	if resp.StatusCode() != 200 {
// 		err = reqErr
// 		return
// 	}

// 	body := resp.String()

// 	// 最后判断是否需要验证码进行登陆
// 	if strings.HasPrefix(body, "false") {
// 		isNeed = false
// 	} else if strings.HasPrefix(body, "true") {
// 		isNeed = true
// 	} else {
// 		err = errors.New("can not get the info")
// 	}

// 	return

// }

// // GetCaptche 获取验证码 (JPEG)
// func (u *SecUser) GetCaptche() (pix []byte, err error) {

// 	resp, err := u.Client.R().
// 		SetQueryParams(map[string]string{
// 			"username": u.Username,
// 			"_":        fmt.Sprint(time.Now().Unix()),
// 		}).
// 		SetHeader("referer", u.AuthUrl).
// 		SetHeader("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8").
// 		Get(u.AuthUrlPerfix + CaptchaPath)

// 	if err != nil {
// 		return
// 	}

// 	pix = resp.Body()

// 	return

// }

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
	actionUrl, err := util.GetSubstringBetweenStringsByRE(body, `id="form" action="`, `"`)
	if err != nil {
		return
	}

	// lt, err := util.GetSubstringBetweenStringsByRE(body, `name="lt" value="`, `"`)
	// if err != nil {
	// 	return
	// }

	execution, err := util.GetSubstringBetweenStringsByRE(body, `name="execution" value="`, `"`)
	if err != nil {
		return
	}

	eventId, err := util.GetSubstringBetweenStringsByRE(body, `name="_eventId" value="`, `"`)
	if err != nil {
		return
	}

	salt, err := util.GetSubstringBetweenStringsByRE(body, `id="salt" value="`, `"`)
	if err != nil {
		return
	}

	// 这个地址需要html解码
	decodeurl := html.UnescapeString(actionUrl)

	// var data = strings.NewReader("username=" + u.Username + "&password=" + u.Password + captchaParam + "&lt=" + lt + "&execution=" + execution + "&_eventId=" + eventId + "&rmShown=" + rmShown)

	dealPassword, err := loginCrypto(u.Password, salt, "1234567890abcdef")
	if err != nil {
		return
	}

	req := u.Client.R().
		SetHeader("referer", u.AuthUrl).
		SetHeader("authority", actionUrl).
		SetFormData(map[string]string{
			"username": u.Username,
			"password": dealPassword,
			// "lt":        lt,
			"execution": execution,
			"_eventId":  eventId,
			// "salt":    salt,
			"rememberMe":  "true", // 一周内免登录 on/off
			"_rememberMe": "on",
		})

	// // 预定.....
	// if len(captcha) > 0 {
	// 	req.SetFormData(map[string]string{
	// 		"captchaResponse": captcha,
	// 	})
	// }

	resp, _ = req.Post(decodeurl)
	if err != nil {
		return
	}

	body = resp.String()

	// log.Println(body)

	// 判断是否有错误
	if strings.Contains(body, "credential.errors") {
		loginErrStr, _ := util.GetSubstringBetweenStringsByRE(body, `credential.errors">`, `</span>`)
		err = errors.New(loginErrStr)
		return
	}

	// 确保账号登陆成功
	// if !u.IsLogged() {
	// 	u.login(captcha)
	// }

	// 获取门户path
	err = u.PerSetPortalPath()

	return
}

// Login 第一层普通登陆
func (u *SecUser) Login() (err error) {
	return u.login("")
}

// // LoginWithCap 第一层验证码登陆
// func (u *SecUser) LoginWithCap(captcha string) (err error) {
// 	return u.login(captcha)
// }

// PortalLogin 第二层门户登陆
func (u *SecUser) PortalLogin() (err error) {

	if len(u.PortalUrlPerfix) == 0 {
		err = errors.New("please login first")
		return
	}

	// 增加重定向次数
	tmpClient := u.Client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

	resp, reqErr := tmpClient.R().
		SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9").
		SetHeader("referer", u.PortalUrlPerfix+PortalIndexPath).
		Get(u.PortalUrlPerfix + PortalLoginPath + "?vpn-0")

	if resp.StatusCode() != http.StatusOK {
		err = reqErr
		return
	}

	// 确保账号登陆成功
	if !u.IsPortalLogged() {
		err = errors.New("fail to login")
	}

	return
}
