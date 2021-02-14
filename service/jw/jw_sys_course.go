package jw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// QueryTermParam 获得课表查询所需参数(学年学期以及院系)
func QueryTermParam(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, ClassSelURL, nil)
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", ClassSelURL)
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

// QueryMajorParam 获得课表查询所需参数(专业)
func QueryMajorParam(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, MajorListURL, nil)

	// 添加请求参数
	q := r.URL.Query()

	q.Add("yx", "09")
	q.Add("nj", "2020")

	r.URL.RawQuery = q.Encode()

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", ClassSelURL)
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

// QueryCourseParam 获得课表查询所需参数(班级)
func QueryCourseParam(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, ClassListURL, nil)

	// 添加请求参数
	q := r.URL.Query()

	q.Add("zy", "0901")
	q.Add("nj", "2020")

	r.URL.RawQuery = q.Encode()

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", ClassSelURL)
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

// QueryCourseTable 查询课表页面
func QueryCourseTable(cookies []*http.Cookie) (string, error) {
	client := &http.Client{}

	data := url.Values{
		"Sel_XNXQ": {"20201"},
		"Submit01": {""},
		"Sel_NJ":   {"2020"},
		"Sel_YX":   {"09"},
		"Sel_ZY":   {"0901"},
		"Sel_BJ":   {"2020090101"},
	}

	r, err := http.NewRequest(http.MethodPost, ClassRptURL, strings.NewReader(data.Encode()))
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
	r.Header.Add("Referer", ClassSelURL)
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

	fmt.Println(bodystr)

	return bodystr, nil
}
