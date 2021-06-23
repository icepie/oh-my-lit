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

	secUser, err := sec.NewSecUser("B19071121", "xxx")
	if err != nil {
		log.Println("实例化用户失败: ", err)
	}

	isNeedcap, err := secUser.IsNeedCaptcha()
	if err != nil {
		log.Println("获取状态信息失败: ", err)
	}

	if isNeedcap {
		log.Println("需要验证码")

		pix, err := secUser.GetCaptche()

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

		err = secUser.LoginWithCap(capp)
		if err != nil {
			log.Fatal("登陆失败: ", err)
		}

	} else {
		log.Println("不需要验证码")
		err = secUser.Login()
		if err != nil {
			log.Fatal("登陆失败: ", err)
		}
	}

	log.Println("登陆成功!")

	if secUser.IsLogged() {

		// 进行门户登陆
		secUser.PortalLogin()

		if secUser.IsPortalLogged() {

			test, err := secUser.GetCurrentMember()
			if err != nil {
				log.Fatal("获取个人信息失败: ", err)
			}

			log.Println("欢迎!", test.Obj.MemberNickname, test.Obj.RoleList[0].RoleName)
			log.Println("上次登陆时间", test.Obj.LastLoginTime)

			t1, err := secUser.GetStudent(secUser.Username)
			if err != nil {
				log.Fatal("查询学生信息失败: ", err)
			}

			log.Println(t1)

			t2, err := secUser.GetStudent("111")
			if err != nil {
				log.Fatal("查询学生信息失败: ", err)
			}

			log.Println(t2)
		}

	}

}
