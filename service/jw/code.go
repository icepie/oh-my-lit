package jw

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetDayJCSel 请求特殊课表查询页面
func GetDayJCSel(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, DayJCSelURL, nil)
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", DayJCSelURL)
	r.Header.Add("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将 gb2312 转换为 utf-8
	bodystr := gb18030Tutf8(string(b))

	fmt.Println(bodystr)

	return bodystr, nil
}
