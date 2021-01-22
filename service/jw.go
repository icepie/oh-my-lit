package service

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

const (
	// LoginURL 登陆地址
	LoginURL = "http://jw.sec.lit.edu.cn/_data/index_LOGIN.aspx"
	// SchoolCode 院校代号
	SchoolCode = "11070"
)

/*
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
func md5s(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// chkpwd 将用户密码进行处理
func chkpwd(username string, password string) string {

	return strings.ToUpper(md5s(username + strings.ToUpper(md5s(password)[0:30]) + SchoolCode)[0:30])
}

// SendLogin 发送登陆表单
func SendLogin() {
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

	fmt.Println(VS)

}
