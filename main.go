package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/icepie/lit-edu-go/conf"
	"github.com/icepie/lit-edu-go/service/jw"
)

func main() {

	// 初始化全局配置
	conf.INIT()

	// 尝试登陆
	cookies, err := jw.SendLogin(conf.ProConf.JW.UserName, conf.ProConf.JW.PassWord)
	if err != nil {
		log.Println(err)
	}

	print(cookies)
}
