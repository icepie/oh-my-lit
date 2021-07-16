package jw

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/icepie/oh-my-lit/client/sec"
	"github.com/icepie/oh-my-lit/client/util"
)

// JwUser Login
func (u *JwUser) Login() (err error) {

	if u.IsBoundSec {
		// 暂时未实现登陆任意帐号
		err = errors.New("please use LoginBySec()")
		return
		//LoginUrl += "?vpn-0"
	}

	client := &http.Client{}

	LoginUrl := u.Url.String() + LoginPath

	req, _ := http.NewRequest(http.MethodGet, LoginUrl, nil)

	// 整上
	for _, cookie := range u.Cookies {
		req.AddCookie(cookie)
	}

	// command, _ := http2curl.GetCurlCommand(req)
	// fmt.Println(command)

	resp1, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp1.Body.Close()

	if len(u.Cookies) == 0 {
		u.Cookies = resp1.Cookies()
	}

	// // 取得 Cookies
	// u.Cookies = resp1.Cookies()

	// 将数据流转换为 []byte
	b, _ := ioutil.ReadAll(resp1.Body)
	// if err != nil {
	// 	return
	// }

	// 将 gb2312 转换为 utf-8
	bodystr := util.GB18030ToUTF8(string(b))

	// log.Println(string(bodystr))

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

	// 不重定向
	// client = &http.Client{
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 		return http.ErrUseLastResponse
	// 	},
	// }

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
		r.Header.Set("referer", LoginUrl)
		r.Header.Set("accept-language", "zh-CN,zh;q=0.9")

	} else {

		r.Header.Set("Host", u.Url.Host)
		r.Header.Set("Proxy-Connection", "keep-alive")
		r.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
		r.Header.Set("Origin", u.Url.String())
		r.Header.Set("Upgrade-Insecure-Requests", "1")
		r.Header.Set("DNT", "1")
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Set("User-Agent", UserAgent)
		r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Header.Set("Referer", LoginUrl)
		r.Header.Set("Accept-Encoding", "gzip, deflate")
		r.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	}

	for _, cookie := range u.Cookies {
		r.AddCookie(cookie)
	}

	// command, _ = http2curl.GetCurlCommand(r)
	// fmt.Println(command)

	// log.Fatal()

	resp, err := client.Do(r)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// log.Println(util.GB18030ToUTF8(string(b)))

	isLogged := u.IsLogged()

	if !isLogged {
		err = errors.New("jw fail to login")
	}

	return

}

func (u *JwUser) LoginBySec() (err error) {

	if !u.IsBoundSec {
		// 暂时未实现登陆任意帐号
		err = errors.New("only for sec user")
		return
		//LoginUrl += "?vpn-0"
	}

	LoginBySecUrl := u.Url.String() + LoginBySecPath + "?vpn-0"

	// if u.IsBoundSec {
	// 	MAINFRMURL += "?vpn-0"
	// }

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, LoginBySecUrl, nil)

	// 整上
	for _, cookie := range u.Cookies {
		req.AddCookie(cookie)
	}

	// command, _ := http2curl.GetCurlCommand(req)
	// fmt.Println(command)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	isLogged := u.IsLogged()

	if !isLogged {
		err = errors.New("jw fail to login")
	}

	return

}

// 是否登陆
func (u *JwUser) IsLogged() (isLogged bool) {

	isLogged = false

	MAINFRMURL := u.Url.String() + MenuPath

	if u.IsBoundSec {
		MAINFRMURL += "?vpn-0"
	}

	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, MAINFRMURL, nil)

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
		r.Header.Set("referer", MAINFRMURL)
		r.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	} else {
		r.Header.Set("User-Agent", UserAgent)
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

	// log.Println(util.GB18030ToUTF8(string(b)))

	// log.Println(r.Cookies())

	// log.Println(MAINFRMURL)

	// command, _ := http2curl.GetCurlCommand(r)
	// fmt.Println(command)

	// 检测是否登陆成功
	if strings.Contains(body, "洛阳理工学院教务") {
		isLogged = true
	}

	return
}
