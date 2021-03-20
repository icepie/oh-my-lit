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
	// BannerURL 菜单地址
	BannerURL = "http://jw.sec.lit.edu.cn/SYS/Main_banner.aspx"
	// MAINFRMURL 主页
	MAINFRMURL = "http://jw.sec.lit.edu.cn/MAINFRM.aspx"
	// ScoreURL 成绩检索地址
	ScoreURL = "http://jw.sec.lit.edu.cn/XSCJ/f_cjdab_rpt.aspx"
	// ClassSelURL 课表查询首页
	ClassSelURL = "http://jw.sec.lit.edu.cn/ZNPK/ClassSel.aspx"
	// MajorListURL 获取专业对应值
	// http://jw.sec.lit.edu.cn/XSXJ/Private/List_NJYXZY.aspx?yx=09&nj=2020
	MajorListURL = "http://jw.sec.lit.edu.cn/XSXJ/Private/List_NJYXZY.aspx"
	// ClassListURL 获取班级对应值
	// http://jw.sec.lit.edu.cn/XSXJ/Private/List_ZYBJ.aspx?zy=0901&nj=2020
	ClassListURL = "http://jw.sec.lit.edu.cn/XSXJ/Private/List_ZYBJ.aspx"
	// ClassRptURL 管理员课表查询地址
	ClassRptURL = "http://jw.sec.lit.edu.cn/ZNPK/ClassSel_rpt.aspx"
	//DayJCSel 隐藏查询地址
	DayJCSelURL = "http://jw.sec.lit.edu.cn/ZNPK/KBFB_DayJCSel.aspx"
	//STUZXJGURL 学生正选结果页面
	STUZXJGURL = "http://jw.sec.lit.edu.cn//wsxk/stu_zxjg_rpt.aspx"
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

func G2U(s string) string {
	return gb18030Tutf8(s)
}

// JWCookies 教务在线曲奇饼
var JWCookies []*http.Cookie

// Test 测试用
func Test() {
	// QueryTermParam(JWCookies)
	// QueryMajorParam(JWCookies)
	// QueryCourseParam(JWCookies)
	// QueryCourseTable(JWCookies)
	// GetBanner(JWCookies)

	GetDayJCSel(JWCookies)

}

// RefreshCookies 刷新教务在线曲奇饼
func RefreshCookies() {
	for {
		var err error
		// 管理人员帐号登陆 SYS
		log.Println("Refreshing jw cookies... ")
		// 最多重试刷新三次
		for i := 0; i < 3; i++ {
			JWCookies, err = SendLogin(conf.ProConf.JW.UserName, conf.ProConf.JW.PassWord, "SYS")

			if err != nil {
				log.Println("Retrying...")
				log.Warningln(err)
			} else {
				log.Println("jw is work fine")
				// 仅在测试中使用
				//Test()
				break
			}

		}
		// 等待下次刷新
		time.Sleep(time.Second * time.Duration(conf.ProConf.JW.RefInt))
	}
}

// IsWork 检查曲奇饼可用性
func IsWork() (bool, error) {
	lflag, err := IsLogged(JWCookies)
	if err != nil || !lflag {
		return false, err
	}
	return true, nil
}
