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

	log.Println("智慧控电测试")

	zhydUser, err := zhyd.NewZhydUser("B19071121", "")
	if err != nil {
		log.Println("实例化用户失败: ", err)
	}

	b, err := zhydUser.IsNeedCaptcha()
	if err != nil {
		log.Println("获取用户信息失败: ", err)
	}

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

	isLogged, err := zhydUser.IsLogged()
	if isLogged {
		de, err := zhydUser.GetDormElectricity()
		if err != nil {
			log.Fatal("获取余电额度失败: ", err)
		}

		log.Println(de)

		ed, err := zhydUser.GetElectricityDetails()
		if err != nil {
			log.Fatal("获取历史用电失败: ", err)
		}

		log.Println(ed)

		cr, err := zhydUser.GetChargeRecords()
		if err != nil {
			log.Fatal("获取充值记录失败: ", err)
		}

		log.Println(cr)

	} else {
		log.Println("似乎未登陆: ", err)
	}

}
