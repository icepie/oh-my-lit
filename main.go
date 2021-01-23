package main

import (
	"fmt"

	"github.com/icepie/lit-edu-go/conf"
	"github.com/icepie/lit-edu-go/server"
	"github.com/icepie/lit-edu-go/service/jw"
)

func main() {

	// 初始化全局配置
	conf.INIT()

	// 尝试登陆, 刷新曲奇饼
	jw.RefreshCookies()

	// 检索 学号为B19071121 2020学年 第一学期 的成绩
	fmt.Println(jw.QueryScoreByStuNum(jw.JWCookies, "B19071122"))

	// 装载路由
	r := server.NewRouter()

	// 运行服务
	r.Run(conf.ProConf.MAIN)
}
