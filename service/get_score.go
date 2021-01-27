package service

import (
	"fmt"
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

	//var Scorel3 Term
	// 去掉空表 (倒数第四个)
	//doc.Find("table").Eq(-4).Remove().End()
	// 成绩藏在第二个表
	table1 := doc.Find("table").Eq(1)
	// 新建一个学期成绩列表
	var newtermList model.TermList
	// 学期个数的计数器
	var Tcount int
	// 查找平均成绩个数, 得出学期的个数
	doc.Find("script").Each(func(index int, tr *goquery.Selection) {
		// 新建一个学期成绩结构
		var newterm model.Term
		// 处理获取到平均成绩: T2.innerHTML='(平均成绩：85.4)
		newterm.AvgScore = exutf8.RuneSubString(tr.Text(), 20, 4)
		// 再扔进列表里
		newtermList = append(newtermList, newterm)
		// 计数器会刷到最终次数
		Tcount = index
	})

	// 逆序循环处理
	for {
		// 循环退出判断
		if Tcount == -1 {
			break
		}
		// 找到该表的id存在的位置
		id := fmt.Sprintf("%s%d", "td#T", Tcount+1)
		//fmt.Println(id)
		// 学期名整上 从doc取得, 因为table1要进行删除操作
		newtermList[Tcount].Term = doc.Find(id).Prev().Text()

		T := table1.Find(id)
		// 找到成绩表所在地方
		T.Prev().ParentsFiltered("tr[style]").NextAllFiltered("tr[style]").Each(func(index int, tr *goquery.Selection) {
			// 新建个成绩结构
			var newscore model.Score
			// td里面包含具体值
			td := tr.Find("td[width]")
			// 以下对号入座
			newscore.Course = td.First().Text()
			newscore.Type = td.Eq(1).Text()
			newscore.Count = td.Eq(2).Text()
			newscore.Score = td.Eq(3).Text()
			newscore.Credit = td.Eq(4).Text()
			// 过滤掉取了空td的情况
			if newscore.Credit != "" {
				// 将成绩添加到成绩表
				newtermList[Tcount].ScoreList = append(newtermList[Tcount].ScoreList, newscore)
			}
			// 加完就删, 逆序处理的核心
			tr.Remove().End()
		})
		// 计数器变化
		Tcount--
	}
	
	// 新建成绩序列化数据
	scoredata := model.ScoreInfo{
		SI: stu,
		TL: newtermList,
	}

	// 调试用
	fmt.Println(scoredata)

	code := e.SUCCESS
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: scoredata,
	}
}
