package jw

import (
	"github.com/axgle/mahonia"
)

const (
	// HostURL 主网址
	HostURL = "jw.sec.lit.edu.cn"
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

func gb2312Tutf8(s string) string {
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
