package jw

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetBanner 获取Banner页面
func GetBanner(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, BannerURL, nil)
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", BannerURL)
	r.Header.Add("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// 将 gb2312 转换为 utf-8
	bodystr := gb18030Tutf8(string(b))

	// 检测是否登陆成功
	if strings.Contains(bodystr, "bakend2") == true {
		return "", errors.New("lit jw can not to login")
	}

	return bodystr, nil
}
