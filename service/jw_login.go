package service

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

const (
	// DefaultURL 首页
	DefaultURL = "http://jw.sec.lit.edu.cn/default.aspx"
	// LoginURL 登陆地址
	LoginURL = "http://jw.sec.lit.edu.cn/_data/index_LOGIN.aspx"
	// MenuURL 菜单地址
	MenuURL = "http://jw.sec.lit.edu.cn/frame/menu.aspx"
	// MAINFRMURL 主页
	MAINFRMURL = "http://jw.sec.lit.edu.cn/MAINFRM.aspx"
	// SchoolCode 院校代号
	SchoolCode = "11070"
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
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// 青果不是utf-8 所以要转换一下
	utf8Body, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(utf8Body)
	if err != nil {
		log.Fatal(err)
	}

	var VS string

	doc.Find("input[name=__VIEWSTATE]").Each(func(i int, s *goquery.Selection) {
		vs, ex := s.Attr("value")
		if ex {
			VS = vs
		}
	})

	return VS, res.Cookies(), err
}

// SendLogin 发送登陆表单
func SendLogin(username string, password string) ([]*http.Cookie, error) {

	vs, cookie, err := getVSAndCookie()

	if err != nil {
		log.Fatal(err)
	}

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

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, LoginURL, strings.NewReader(data.Encode()))
	r.Header.Add("Host", "jw.sec.lit.edu.cn")
	r.Header.Add("Proxy-Connection", "keep-alive")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Add("Origin", "http://jw.sec.lit.edu.cn")
	r.Header.Add("Upgrade-Insecure-Requests", "1")
	r.Header.Add("DNT", "1")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Header.Add("Referer", LoginURL)
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(r.URL, cookie)

	resp, _ := client.Do(r)

	fmt.Println(resp.Cookies())

	fmt.Println(resp.Request)

	r, _ = http.NewRequest(http.MethodGet, MenuURL, nil)

	jar, _ = cookiejar.New(nil)
	jar.SetCookies(r.URL, cookie)

	r.Header.Add("Host", "jw.sec.lit.edu.cn")
	r.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Referer", MenuURL)
	r.Header.Add("Upgrade-Insecure-Requests", "1")

	resp, _ = client.Do(r)

	fmt.Println(resp.Cookies())

	// 青果不是utf-8 所以要转换一下
	utf8Body, err := iconv.NewReader(resp.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadAll(utf8Body)
	resp.Body.Close()

	fmt.Println(string(b))

	return cookie, err
}
