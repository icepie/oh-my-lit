package service

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

// getNewCookie 获取新 Cookie, 主要是新 ASP.NET_SessionId
func getNewCookie() ([]*http.Cookie, error) {
	res, err := http.Get(DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	return res.Cookies(), err
}

// getViewState 获取 _VIEWSTATE
func getViewState() (string, error) {
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

	return VS, err
}

// SendLogin 发送登陆表单
func SendLogin(username string, password string) ([]*http.Cookie, error) {

	cookie, err := getNewCookie()

	if err != nil {
		log.Fatal(err)
	}

	vs, err := getViewState()
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
	//r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(r.URL, cookie)

	resp, _ := client.Do(r)

	fmt.Println(resp.Request)

	//clt := http.Client{Transport: nil, Jar: jar}

	// clt.PostForm(url.String(), map[string][]string{
	// 	"__VIEWSTATE":             {vs},
	// 	"Sel_Type":                {"SYS"}, // only for SYS
	// 	"txt_sdsdfdsfryuiighgdf":  {username},
	// 	"txt_dsfdgtjhjuixssdsdf":  {},
	// 	"txt_sftfgtrefjdndcfgerg": {},
	// 	"typeName":                {},
	// 	"sdfdfdhgwerewt":          {chkpwd(username, password)},
	// 	"cxfdsfdshjhjlk":          {},
	// })

	r, _ = http.NewRequest(http.MethodGet, MenuURL, nil)
	jar, _ = cookiejar.New(nil)
	jar.SetCookies(r.URL, cookie)

	resp, _ = client.Do(r)

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

// ISLogin 判是登陆成功性 for test
func ISLogin(username string, pasword string) {
	cookie, err := SendLogin(username, pasword)
	if err != nil {
		log.Fatal(err)
	}

	url, err := url.Parse("http://jw.sec.lit.edu.cn/frame/menu.aspx")

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	jar.SetCookies(url, cookie) //这里的cookies是[]*http.Cookie

	//fmt.Println(cookie)

	clt := http.Client{Transport: nil, Jar: jar}

	resp, _ := clt.Get(url.String())

	utf8Body, err := iconv.NewReader(resp.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadAll(utf8Body)
	resp.Body.Close()

	fmt.Println(string(b))

}
