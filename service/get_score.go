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

// GetScoreService 获取成绩服务结构
type GetScoreService struct {
	StuID string `json:"stuid" binding:"required"`
}

// GetScore 根据 StuID 获取
func (service *GetScoreService) GetScore() model.Response {
	var stu model.Stu

	body, err := jw.QueryScoreByStuNum(jw.JWCookies, service.StuID)
	if err != nil {
		log.Println(err)
		code := e.ERROR
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println(err)
		code := e.ERROR
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// // Find the review items
	// // Find each table
	// doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
	// 	tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
	// 		rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
	// 			headings = append(headings, tableheading.Text())
	// 		})
	// 		rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
	// 			row = append(row, tablecell.Text())
	// 		})
	// 		rows = append(rows, row)
	// 		row = nil
	// 	})
	// })

	// 经济与管理学院
	// 本科
	// 四年
	// 2019年09月
	// B19071121
	// 工商管理
	// B190701
	// 2023年03月
	// XXX

	// 解析学生信息
	doc.Find("tbody").First().Each(func(i int, tbody *goquery.Selection) {
		tbody.Find("tr").First().Each(func(i int, tr *goquery.Selection) {
			tr.Find("td").First().Each(func(i int, td *goquery.Selection) {
				stu.Faculty = exutf8.RuneSubString(td.Text(), 7, 20)
			})
			tr.Find("td").Eq(1).Each(func(i int, td *goquery.Selection) {
				stu.Degree = exutf8.RuneSubString(td.Text(), 5, 10)
			})
			tr.Find("td").Eq(2).Each(func(i int, td *goquery.Selection) {
				stu.EduSys = exutf8.RuneSubString(td.Text(), 3, 10)
			})
			tr.Find("td").Eq(3).Each(func(i int, td *goquery.Selection) {
				stu.AdmTime = exutf8.RuneSubString(td.Text(), 5, 10)
			})
			tr.Find("td").Eq(4).Each(func(i int, td *goquery.Selection) {
				stu.ID = exutf8.RuneSubString(td.Text(), 3, 10)
			})
		})
		tbody.Find("tr").Eq(1).Each(func(i int, tr *goquery.Selection) {
			tr.Find("td").First().Each(func(i int, td *goquery.Selection) {
				stu.Major = exutf8.RuneSubString(td.Text(), 8, 20)
			})
			tr.Find("td").Eq(1).Each(func(i int, td *goquery.Selection) {
				stu.Class = exutf8.RuneSubString(td.Text(), 5, 10)
			})
			tr.Find("td").Eq(2).Each(func(i int, td *goquery.Selection) {
				stu.GraTime = exutf8.RuneSubString(td.Text(), 5, 10)
			})
			tr.Find("td").Eq(3).Each(func(i int, td *goquery.Selection) {
				stu.Name = exutf8.RuneSubString(td.Text(), 3, 10)
			})
		})
	})

	code := e.SUCCESS
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   stu,
	}
}
