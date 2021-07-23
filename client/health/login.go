package health

import (
	"encoding/json"
	"errors"
	"log"
)

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
		log.Println(err)
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

// IsLogged 判断是否登录
func (u *HealthUser) IsLogged() bool {
	_, err := u.GetLastRecord()
	return err == nil
}
