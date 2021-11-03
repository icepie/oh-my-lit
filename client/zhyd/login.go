package zhyd

import (
	"errors"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

func (u *ZhydUser) IsLogged() bool {

	resp, _ := u.Client.R().
		Get(ZhydHostUrl)

	return !strings.Contains(resp.String(), "统一身份认证")
}

// // IsNeedCaptcha 判断是否需要验证码登陆
// func (u *ZhydUser) IsNeedCaptcha() (isNeed bool, err error) {

// 	resp, reqErr := u.Client.R().
// 		SetQueryParams(map[string]string{
// 			"username": u.Username,
// 			"_":        fmt.Sprint(time.Now().Unix()),
// 		}).
// 		Get(NeedCaptchaUrl)

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
// func (u *ZhydUser) GetCaptche() (pix []byte, err error) {

// 	resp, err := u.Client.R().
// 		SetQueryParams(map[string]string{
// 			"username": u.Username,
// 			"_":        fmt.Sprint(time.Now().Unix()),
// 		}).
// 		SetHeader("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8").
// 		Get(CaptchaUrl)

// 	if err != nil {
// 		return
// 	}

// 	pix = resp.Body()

// 	return
// }

// 登陆凑合用
func (u *ZhydUser) login(captcha string) (err error) {

	if len(u.Username) == 0 {
		err = errors.New("empty username")
		return
	}

	if len(u.Password) == 0 {
		err = errors.New("empty password")
		return
	}

	// 禁止重定向
	// tmpClient := u.Client.SetRedirectPolicy(resty.NoRedirectPolicy())
	// u.Client.SetRedirectPolicy(resty.NoRedirectPolicy())

	resp, err := u.Client.R().
		Get(LoginUrl)
	if err != nil {
		return
	}

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
	// decodeurl := html.UnescapeString(actionUrl)

	// var data = strings.NewReader("username=" + u.Username + "&password=" + u.Password + captchaParam + "&lt=" + lt + "&execution=" + execution + "&_eventId=" + eventId + "&rmShown=" + rmShown)

	dealPassword, err := loginCrypto(u.Password, salt, "1234567890abcdef")
	if err != nil {
		return
	}

	req := u.Client.R().
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

	resp, err = req.Post(LoginUrl)
	if err != nil {
		return
	}

	body = resp.String()

	// 判断是否有错误
	if strings.Contains(resp.String(), "credential.errors") {
		loginErrStr, _ := util.GetSubstringBetweenStringsByRE(body, `credential.errors">`, `</span>`)
		err = errors.New(loginErrStr)
		return
	}

	return
}

// Login 普通登陆
func (u *ZhydUser) Login() (err error) {
	// 操蛋玩意,多登陆几次
	for i := 0; i <= 2; i++ {
		err = u.login("")
		if err == nil {
			return
		}
	}
	return
}

// // LoginWithCap 验证码登陆
// func (u *ZhydUser) LoginWithCap(captcha string) error {
// 	return u.login(captcha)
// }
