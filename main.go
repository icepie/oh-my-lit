package main

import (
	"github.com/icepie/lit-edu-go/conf"
	"github.com/icepie/lit-edu-go/server"
	"github.com/icepie/lit-edu-go/service/jw"
)

func main() {

	// 尝试登陆, 刷新曲奇饼
	go jw.RefreshCookies()

	// 测试用
	// QueryTermParam(JWCookies)
	// QueryMajorParam(JWCookies)
	// QueryClassParam(JWCookies)
	// QueryClassTable(JWCookies)

	// 装载路由
	r := server.ApiRouter()

	// 运行服务
	r.Run(conf.MAIN)
}
