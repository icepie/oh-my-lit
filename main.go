package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/icepie/lit-edu-go/conf"
	"github.com/icepie/lit-edu-go/service/jw"
)

// JWCookies 教务在线曲奇饼
var JWCookies []*http.Cookie

// JWUserName 教务在线用户名
var JWUserName string

// JWPassWord 教务在线密码
var JWPassWord string

// RefreshCookies 刷新教务在线曲奇饼
func RefreshCookies() {
	var err error
	JWCookies, err = jw.SendLogin(JWUserName, JWPassWord)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	// 初始化全局配置
	conf.INIT()
	JWUserName = conf.ProConf.JW.UserName
	JWPassWord = conf.ProConf.JW.PassWord

	// 尝试登陆, 刷新曲奇饼
	RefreshCookies()

	// 检索 学号为B19071121 2020学年 第一学期 的成绩
	fmt.Println(jw.QueryScoreByStuNum(JWCookies, "B19071122"))
}
