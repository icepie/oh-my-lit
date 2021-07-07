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

	if len(os.Args) != 2 {
		return
	}

	secUser, err := sec.NewSecUser("B19071121", os.Args[1])
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

			t2, err := secUser.GetClassmates(secUser.Username)
			if err != nil {
				log.Fatal("查询同班同学列表失败: ", err)
			}

			log.Println(t2)

			t3, err := secUser.GetClassmatesDetail(secUser.Username)
			if err != nil {
				log.Fatal("查询同班同学信息失败: ", err)
			}

			log.Println(t3)

			t4, err := secUser.GetOneCardBalance(secUser.Username)
			if err != nil {
				log.Fatal("查询一卡通余额失败: ", err)
			}

			log.Println(t4)

			t5, err := secUser.GetOneCardChargeRecords(secUser.Username, 1, 10)
			if err != nil {
				log.Fatal("查询一卡通充值记录失败: ", err)
			}

			log.Println(t5)

			t6, err := secUser.GetOneCardConsumeRecords(secUser.Username, 1, 10)
			if err != nil {
				log.Fatal("查询一卡通充值记录失败: ", err)
			}

			log.Println(t6)

			t7, err := secUser.GetExamArrangements(secUser.Username, 2020, 1)
			if err != nil {
				log.Fatal("查询考试安排失败: ", err)
			}

			log.Println(t7)

			t8, err := secUser.GetWeekCourses(secUser.Username, "2021-05-07", 1)
			if err != nil {
				log.Fatal("查询周课表失败: ", err)
			}

			log.Println(t8)
		}

	}
}
