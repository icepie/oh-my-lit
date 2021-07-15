package jw

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/icepie/oh-my-lit/client/sec"
	"github.com/icepie/oh-my-lit/client/util"
)

// JwUser reset MianURL
func (u *JwUser) Login() (err error) {

	client := &http.Client{}

	LoginUrl := u.Url.String() + LoginPath

	if u.IsBoundSec {
		LoginUrl += "?vpn-0"
	}

	req, _ := http.NewRequest(http.MethodGet, LoginUrl, nil)

	for _, cookie := range u.Cookies {
		req.AddCookie(cookie)
	}

	log.Println(LoginUrl)

	resp1, err := client.Do(req)
	if err != nil {
		return
	}

	if len(u.Cookies) == 0 {
		u.Cookies = resp1.Cookies()
	}

	log.Println(LoginUrl)

	defer resp1.Body.Close()

	// // 取得 Cookies
	// u.Cookies = resp1.Cookies()

	// 将数据流转换为 []byte
	b, _ := ioutil.ReadAll(resp1.Body)
	// if err != nil {
	// 	return
	// }

	log.Println(LoginUrl)

	// 将 gb2312 转换为 utf-8
	bodystr := util.GB18030ToUTF8(string(b))

	log.Println(string(bodystr))

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodystr))
	if err != nil {
		return
	}

	// 拿到参数
	vs, isExist := doc.Find("input[name=__VIEWSTATE]").First().Attr("value")
	if !isExist {
		err = errors.New("vs is no exist")
		return
	}

	log.Println(vs)

	data := url.Values{
		"__VIEWSTATE":             {vs},
		"Sel_Type":                {u.SelType}, // SYS etc..
		"txt_sdsdfdsfryuiighgdf":  {u.Username},
		"txt_dsfdgtjhjuixssdsdf":  {},
		"txt_sftfgtrefjdndcfgerg": {},
		"typeName":                {},
		"sdfdfdhgwerewt":          {chkpwd(u.Username, u.Password)},
		"cxfdsfdshjhjlk":          {},
	}

	r, err := http.NewRequest(http.MethodPost, LoginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}

	if u.IsBoundSec {
		r.Header.Set("authority", u.Url.Host)
		r.Header.Set("cache-control", "max-age=0")
		r.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
		r.Header.Set("sec-ch-ua-mobile", "?0")
		r.Header.Set("dnt", "1")
		r.Header.Set("upgrade-insecure-requests", "1")
		r.Header.Set("user-agent", sec.UA)
		r.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Header.Set("sec-fetch-site", "same-origin")
		r.Header.Set("sec-fetch-mode", "navigate")
		r.Header.Set("sec-fetch-user", "?1")
		r.Header.Set("sec-fetch-dest", "document")
		r.Header.Set("referer", "https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?vpn-0&service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F")
		r.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	} else {

		r.Header.Add("Host", u.Url.Host)
		r.Header.Add("Proxy-Connection", "keep-alive")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		r.Header.Add("Origin", u.Url.String())
		r.Header.Add("Upgrade-Insecure-Requests", "1")
		r.Header.Add("DNT", "1")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("User-Agent", UserAgent)
		r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Header.Add("Referer", LoginUrl)
		r.Header.Add("Accept-Encoding", "gzip, deflate")
		r.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	}

	for _, cookie := range u.Cookies {
		r.AddCookie(cookie)
	}

	resp, err := client.Do(r)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	log.Println(util.GB18030ToUTF8(string(b)))

	isLogged := u.IsLogged()

	if !isLogged {
		err = errors.New("jw fail to login")
	}

	return

}

// 是否登陆
func (u *JwUser) IsLogged() (isLogged bool) {

	MenuURL := u.Url.String() + MenuPath

	if u.IsBoundSec {
		MenuURL += "?vpn-0"
	}

	client := &http.Client{}

	isLogged = false

	r, _ := http.NewRequest(http.MethodGet, MenuURL, nil)

	for _, cookie := range u.Cookies {
		r.AddCookie(cookie)
	}

	if u.IsBoundSec {
		r.Header.Set("authority", u.Url.Host)
		r.Header.Set("cache-control", "max-age=0")
		r.Header.Set("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
		r.Header.Set("sec-ch-ua-mobile", "?0")
		r.Header.Set("dnt", "1")
		r.Header.Set("upgrade-insecure-requests", "1")
		r.Header.Set("user-agent", sec.UA)
		r.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Header.Set("sec-fetch-site", "same-origin")
		r.Header.Set("sec-fetch-mode", "navigate")
		r.Header.Set("sec-fetch-user", "?1")
		r.Header.Set("sec-fetch-dest", "document")
		r.Header.Set("referer", "https://sec.lit.edu.cn/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?vpn-0&service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F")
		r.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	} else {
		r.Header.Add("Host", u.Url.Host)
		r.Header.Add("User-Agent", UserAgent)
		r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
		r.Header.Add("Accept-Encoding", "gzip, deflate")
		r.Header.Add("Connection", "keep-alive")
		r.Header.Add("Referer", MenuURL)
		r.Header.Add("Upgrade-Insecure-Requests", "1")
	}

	resp, err := client.Do(r)
	if err != nil {
		return
	}

	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body := util.GB18030ToUTF8(string(b))

	log.Println(util.GB18030ToUTF8(string(b)))

	log.Println(r.Cookies())

	// 检测是否登陆成功
	if strings.Contains(body, "洛阳理工学院教务网站") {
		isLogged = true
	}

	return
}
