package jw

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
)

func md5s(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// chkpwd 将用户密码进行处理
func chkpwd(username string, password string) string {

	/* JavaScript
	function chkpwd(obj)
	{
		var schoolcode="11070";
		var yhm=document.all.txt_sdsdfdsfryuiighgdf.value;
		if(obj.value!="")
		{
			if(document.all.Sel_Type.value=="ADM")
				yhm=yhm.toUpperCase();
			var s=md5(yhm+md5(obj.value).substring(0,30).toUpperCase()+schoolcode).substring(0,30).toUpperCase();
			document.all.sdfdfdhgwerewt.value=s;
		}
		else
		{
			document.all.sdfdfdhgwerewt.value=obj.value;
		}
	}
	*/

	return strings.ToUpper(md5s(username + strings.ToUpper(md5s(password)[0:30]) + SchoolCode)[0:30])
}

// getVSAndCookie 获取 _VIEWSTAT和 Cookie
func getVSAndCookie() (string, []*http.Cookie, error) {
	res, err := http.Get(LoginURL)
	if err != nil {
		var t []*http.Cookie
		return "", t, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("status code error: %d %s", res.StatusCode, res.Status)
		var t []*http.Cookie
		return "", t, errors.New("status code error")
	}

	// // 青果不是utf-8 所以要转换一下
	// utf8Body, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
	// if err != nil {
	// 	log.Println(err)
	// }
	// iconv 依赖于 libiconv 对 windows 不友好

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		var t []*http.Cookie
		return "", t, err
	}

	// 将 gb2312 转换为 utf-8
	bodystr := gb18030Tutf8(string(b))

	// println(string(bodystr))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodystr))
	if err != nil {
		var t []*http.Cookie
		return "", t, err
	}

	var VS string

	doc.Find("input[name=__VIEWSTATE]").Each(func(i int, s *goquery.Selection) {
		vs, ex := s.Attr("value")
		if ex {
			VS = vs
		}
	})

	return VS, res.Cookies(), nil
}

func isLogged(cookies []*http.Cookie) (bool, error) {
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, MenuURL, nil)
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
	r.Header.Add("Host", HostURL)
	r.Header.Add("User-Agent", UserAgent)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", MenuURL)
	r.Header.Add("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(r)
	if err != nil {
		return false, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	// 检测是否登陆成功
	if strings.Contains(gb18030Tutf8(string(b)), "bakend2") == true {
		return false, errors.New("lit jw can not to login")
	}

	return true, nil
}

// SendLogin 发送登陆表单
func SendLogin(username string, password string) ([]*http.Cookie, error) {

	vs, cookies, err := getVSAndCookie()

	client := &http.Client{}

	data := url.Values{
		"__VIEWSTATE":             {vs},
		"Sel_Type":                {"SYS"}, // only for SYS
		"txt_sdsdfdsfryuiighgdf":  {username},
		"txt_dsfdgtjhjuixssdsdf":  {},
		"txt_sftfgtrefjdndcfgerg": {},
		"typeName":                {},
		"sdfdfdhgwerewt":          {chkpwd(username, password)},
		"cxfdsfdshjhjlk":          {},
	}

	r, err := http.NewRequest(http.MethodPost, LoginURL, strings.NewReader(data.Encode()))
	if err != nil {
		var t []*http.Cookie
		return t, err
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
		var t []*http.Cookie
		return t, err
	}

	defer resp.Body.Close()

	lflag, err := isLogged(cookies)
	if err != nil || lflag == false {
		var t []*http.Cookie
		return t, err
	}

	return cookies, err
}
