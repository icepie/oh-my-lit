package sec

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetHomeParam 获取主页参数, 用于检测第一层登陆的成功性
func (u *SecUser) GetHomeParam() (rte HomeParam, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", GetHomeParamUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.Cookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("authority", AuthorityUrl)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", UA)
	req.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://sec.lit.edu.cn/frontend_static/frontend/login/index.html")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyText, &rte)
	if err != nil {
		return
	}

	// 接口错误解析
	if rte.Code != 0 {
		err = errors.New(rte.Msg)
	}

	return

}
