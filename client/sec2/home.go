package sec2

import (
	"encoding/json"
	"errors"
)

// GetHomeParam 获取主页参数, 用于检测第一层登陆的成功性
func (u *SecUser) GetHomeParam() (rte HomeParam, err error) {

	resp, _ := u.Client.R().
		SetHeader("accept", "application/json, text/plain, */*").
		SetHeader("referer", u.AuthUrl).
		SetHeader("referer", SecUrl+"/frontend_static/frontend/login/index.html").
		Get(GetHomeParamUrl)

	err = json.Unmarshal(resp.Body(), &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if rte.Code != 0 {
		err = errors.New(rte.Msg)
	}

	return

}
