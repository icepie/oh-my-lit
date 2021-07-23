package main

import (
	"log"
	"os"

	"github.com/icepie/oh-my-lit/client/health"
)

func main() {

	// 实例化用户 (链式调用)
	user := health.NewHealthUser().
		SetUsename("B19071121").
		SetPassword(os.Args[1])

	// 设置http客户端参数 (链式调用)
	user.Client.SetProxy("http://proxy.xxxxxxx.cn:12333") //  .SetDebug(true)

	err := user.Login()
	if err != nil {
		log.Fatalln("登陆失败:", err)
	}

	if user.IsLogged() {
		log.Println("你好", user.UserInfo.Name)

		rte, err := user.GetLastRecord()
		if err != nil {
			log.Fatalln("获取上次上报记录失败:", err)
		}

		log.Println("最近上报时间", rte.CreateTime)

		log.Println(rte.Temperature, rte.TemperatureTwo, rte.TemperatureThree)

		isAllReported, err := user.IsReportedToday(0)
		if err != nil {
			log.Fatalln("判断今日上报状态失败:", err)
		}

		if isAllReported {
			log.Println("All Done!")
		}

	} else {
		log.Println("登录失败")
	}

}
