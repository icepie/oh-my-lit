package health

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetToken 根据 Cookies 获取Token信息
func (u *HealthUser) GetToken() (ret UserInfo, err error) {

	resp, err := u.Client.R().
		SetResult(Result{}).
		Get(GetTokenUrl)
	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
		return
	}

	byteData, _ := json.Marshal(r.Data)
	err = json.Unmarshal(byteData, &u.UserInfo)
	if err != nil {
		return
	}

	ret = u.UserInfo
	return
}

// Login 登录健康平台
func (u *HealthUser) Login() (err error) {

	if len(u.Username) == 0 {
		err = errors.New("empty username")
		return
	}

	if len(u.Password) == 0 {
		err = errors.New("empty password")
		return
	}

	// 预先设置cookies
	u.PerSetCooikes()

	resp, err := u.Client.R().
		SetBody(LoginParam{CardNo: u.Username, Password: getSha256(u.Password)}).
		SetResult(Result{}).
		Post(LoginUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
		return
	}

	byteData, _ := json.Marshal(r.Data)
	err = json.Unmarshal(byteData, &u.UserInfo)
	if err != nil {
		return
	}

	u.Client.SetHeader("token", u.UserInfo.Token)

	return
}

// LoginWithSec 健康平台使用统一登陆
func (u *HealthUser) LoginWithSec() (err error) {

	if len(u.Username) == 0 {
		err = errors.New("empty username")
		return
	}

	if len(u.SecPassword) == 0 {
		err = errors.New("empty sec password")
		return
	}

	// 禁止重定向
	// tmpClient := u.Client.SetRedirectPolicy(resty.NoRedirectPolicy())
	// u.Client.SetRedirectPolicy(resty.NoRedirectPolicy())

	resp, err := u.Client.R().
		Get(SecLoginUrl)
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

	dealPassword, err := loginCrypto(u.SecPassword, salt, "1234567890abcdef")
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

	resp, err = req.Post(SecLoginUrl)
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

	mainUrl, _ := url.Parse(MianUrl)

	// 从Cookies获取token
	Cookies := u.Client.GetClient().Jar.Cookies(mainUrl)

	for _, cookie := range Cookies {
		if cookie.Name == "token" {
			u.Client.SetHeader("token", cookie.Value)
			break
		}
	}

	// 补全登陆信息
	ui, err := u.GetToken()
	if err != nil {
		return
	}

	u.UserInfo = ui

	return
}

// IsLogged 判断是否登录
func (u *HealthUser) IsLogged() bool {
	_, err := u.GetLastRecord()
	return err == nil
}
