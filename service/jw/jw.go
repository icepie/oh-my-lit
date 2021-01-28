package jw

import (
	"net/http"
	"time"

	"github.com/axgle/mahonia"
	"github.com/icepie/lit-edu-go/conf"
	log "github.com/sirupsen/logrus"
)

const (
	// HostURL 主网址
	HostURL = "jw.sec.lit.edu.cn" // 120.194.42.205:9001
	// DefaultURL 首页
	DefaultURL = "http://jw.sec.lit.edu.cn/default.aspx"
	// LoginURL 登陆地址
	LoginURL = "http://jw.sec.lit.edu.cn/_data/index_LOGIN.aspx"
	// MenuURL 菜单地址
	MenuURL = "http://jw.sec.lit.edu.cn/frame/menu.aspx"
	// MAINFRMURL 主页
	MAINFRMURL = "http://jw.sec.lit.edu.cn/MAINFRM.aspx"
	// ScoreURL 成绩检索地址
	ScoreURL = "http://jw.sec.lit.edu.cn/XSCJ/f_cjdab_rpt.aspx"
	// UserAgent UA
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
	// SchoolCode 院校代号
	SchoolCode = "11070"
)

func gb18030Tutf8(s string) string {
	src := mahonia.NewDecoder("gb18030")
	res := src.ConvertString(s)
	tag := mahonia.NewDecoder("utf-8")

	_, cdata, err := tag.Translate([]byte(res), true)
	if err != nil {
		return ""
	}

	result := string(cdata)

	return result
}

// JWCookies 教务在线曲奇饼
var JWCookies []*http.Cookie

// RefreshCookies 刷新教务在线曲奇饼
func RefreshCookies() {
	for {
		var err error
		// 管理人员帐号登陆 SYS
		JWCookies, err = SendLogin(conf.ProConf.JW.UserName, conf.ProConf.JW.PassWord, "SYS")
		if err != nil {
			log.Warningln(err)
		} else {
			log.Println("JW is work fine")
		}
		// 等待下次刷新
		time.Sleep(time.Second * time.Duration(conf.ProConf.JW.RefInt))
	}
}

// IsWork 检查曲奇饼可用性
func IsWork() (bool, error) {
	lflag, err := IsLogged(JWCookies)
	if err != nil || lflag == false {
		return false, err
	}
	return true, nil
}
