package service

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/service/jw"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
)

// GetBaseInfoService 获取学生基本信息服务结构
type GetBaseInfoService struct {
	User model.StuAccount
}

// GetBaseInfo 根据 StuID 获取学生基本信息
func (service *GetBaseInfoService) GetBaseInfo() model.Response {

	// // 开启教务密码验证的情况
	// // if conf.ProConf.JWAuth {
	// if service.User.PassWord == "" {
	// 	code := e.Error
	// 	return model.Response{
	// 		Status: code,
	// 		Msg:    e.GetMsg(code),
	// 		Error:  "please enter the correct password",
	// 	}
	// }

	// _, err := jw.SendLogin(service.User.StuID, service.User.PassWord, "STU")
	// if err != nil {
	// 	log.Warningln(err)
	// 	code := e.Error
	// 	return model.Response{
	// 		Status: code,
	// 		Msg:    e.GetMsg(code),
	// 		Error:  err.Error(),
	// 	}
	// }
	// // }

	body, err := jw.QueryScoreByStuNum(jw.JWCookies, service.User.StuID)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 新建学生信息结构变量
	var stu model.Stu
	// 解析学生信息
	table := doc.Find("table").First()
	// 都在第一个table里啦
	table.Find("tbody").First().Each(func(i int, tbody *goquery.Selection) {
		tbody.Find("tr").First().Each(func(i int, tr *goquery.Selection) {
			stu.Faculty = exutf8.RuneSubString(tr.Find("td").First().Text(), 7, 20)
			stu.Degree = exutf8.RuneSubString(tr.Find("td").Eq(1).Text(), 5, 10)
			stu.EduSys = exutf8.RuneSubString(tr.Find("td").Eq(2).Text(), 3, 10)
			stu.AdmTime = exutf8.RuneSubString(tr.Find("td").Eq(3).Text(), 5, 10)
			stu.ID = exutf8.RuneSubString(tr.Find("td").Eq(4).Text(), 3, 10)
		})
		tbody.Find("tr").Eq(1).Each(func(i int, tr *goquery.Selection) {
			stu.Major = exutf8.RuneSubString(tr.Find("td").First().Text(), 8, 20)
			stu.Class = exutf8.RuneSubString(tr.Find("td").Eq(1).Text(), 5, 10)
			stu.GraTime = exutf8.RuneSubString(tr.Find("td").Eq(2).Text(), 5, 10)
			stu.Name = exutf8.RuneSubString(tr.Find("td").Eq(3).Text(), 3, 10)
		})
	})

	// 调试用
	log.Info(stu)

	code := e.Success
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   stu,
	}
}
