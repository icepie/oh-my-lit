package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/icepie/oh-my-lit/client/zhyd"
)

func main() {

	log.Println("智慧门户测试")

	zhydUser, err := zhyd.NewZhydUser("B19071121", "")
	if err != nil {
		log.Println("实例化用户失败: ", err)
	}

	b, err := zhydUser.IsNeedCaptcha()
	if err != nil {
		log.Println("获取用户信息失败: ", err)
	}

	log.Println(zhydUser.Cookies)

	if b {
		log.Println("需要验证码")
		pix, err := zhydUser.GetCaptche()

		if err != nil {
			log.Println("获取验证码失败: ", err)
		}

		out, err := os.Create("./captcha.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, bytes.NewReader(pix))
		if err != nil {
			log.Fatal(err)
		}

		var capp string
		fmt.Print("验证码(./captcha.jpeg): ")
		fmt.Scanf("%s", &capp)

		err = zhydUser.LoginWithCap(capp)
		if err != nil {
			log.Fatal("登陆失败: ", err)
		}

	}

	err = zhydUser.Login()
	if err != nil {
		log.Fatal("登陆失败: ", err)
	}

	zhydUser.GetDormElectricity()

}
