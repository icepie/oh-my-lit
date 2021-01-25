package jw

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// QueryScoreByStuNum 通过学号查询学生成绩
func QueryScoreByStuNum(cookies []*http.Cookie, stuid string) (string, error) {
	client := &http.Client{}

	data := url.Values{
		"sel_xnxq:": {"20190"}, // 学年学期标签好像坏掉了
		"sel_yx":    {"05"},    // only for SYS
		"ChkXH":     {"on"},
		"txtXH":     {stuid},
		"mrxsj":     {},
		"ybysj":     {},
		"mbysjt":    {},
		"radCx":     {"0"},
	}

	r, err := http.NewRequest(http.MethodPost, ScoreURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	r.Header.Add("Host", HostURL)
	r.Header.Add("Proxy-Connection", "keep-alive")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Add("Origin", "http://"+HostURL)
	r.Header.Add("Upgrade-Insecure-Requests", "1")
	r.Header.Add("DNT", "1")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Header.Add("Referer", LoginURL)
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

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

	// 检测是否查询成功
	if strings.Contains(bodystr, stuid) == false {
		return "", errors.New("score info can not be found with " + stuid)
	}

	// 检测登陆是否失效
	if strings.Contains(bodystr, "bakend2") == true {
		return "", errors.New("lit jw can not to login")
	}
	return bodystr, nil
}
