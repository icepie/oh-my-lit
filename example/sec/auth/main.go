package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/icepie/oh-my-lit/client/sec"
)

func main() {

	log.Println("智慧门户测试")

	secUser, err := sec.NewSecUser("B19071121", "xxxxxx")
	if err != nil {
		log.Println("实例化用户失败:", err)
	}

	isNeedcap, err := secUser.IsNeedCaptcha()
	if err != nil {
		log.Println("获取状态信息失败:", err)
	}

	if isNeedcap {
		log.Println("需要验证码")

		pix, err := secUser.GetCaptche()

		if err != nil {
			log.Println("获取验证码失败:", err)
		}

		out, err := os.Create("./xxx1.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, bytes.NewReader(pix))
		if err != nil {
			log.Fatal(err)
		}

		var capp string
		fmt.Scanln("请输入验证码: ", &capp)

		err = secUser.LoginWithCap(capp)
		if err != nil {
			log.Println("登陆失败:", err)
		}

	} else {
		log.Println("不需要验证码")
		err = secUser.Login()
		if err != nil {
			log.Println("登陆失败:", err)
		}
	}

	err = secUser.Login()
	if err != nil {
		log.Println("登陆失败:", err)
	}

	log.Println("登陆成功!")

}
