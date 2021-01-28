package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
)

func main() {

	//var headings, row []string
	//var rows [][]string
	f, e := os.Open("jw-sys-full-score.html")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
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

	var StuInfo string

	doc.Find("tbody").First().Each(func(i int, tbody *goquery.Selection) {
		tbody.Find("tr").First().Each(func(i int, tr *goquery.Selection) {
			tr.Find("td").First().Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 7, 20)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(1).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 5, 10)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(2).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 3, 10)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(3).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 5, 10)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(4).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 3, 10)
				fmt.Println(StuInfo)
			})
		})
		tbody.Find("tr").Eq(1).Each(func(i int, tr *goquery.Selection) {
			tr.Find("td").First().Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 8, 20)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(1).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 5, 10)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(2).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 5, 10)
				fmt.Println(StuInfo)
			})
			tr.Find("td").Eq(3).Each(func(i int, td *goquery.Selection) {
				StuInfo = exutf8.RuneSubString(td.Text(), 3, 10)
				fmt.Println(StuInfo)
			})
		})
	})

	// Score 成绩结构
	type Score struct {
		Course string
		Type   string
		Count  string
		Score  string
		Credit string
	}

	// Term 学期成绩列表结构
	type Term struct {
		Term      string
		ScoreList []Score
		AvgScore  string
	}

	type TermList []Term

	// 新建一个学期成绩列表
	var newtermList TermList
	// 学期个数的计数器
	var Tcount int
	// 查找平均成绩个数, 得出学期的个数
	doc.Find("script").Each(func(index int, tr *goquery.Selection) {
		// 新建一个学期成绩结构
		var newterm Term
		// 处理获取到平均成绩: T2.innerHTML='(平均成绩：85.4)
		newterm.AvgScore = exutf8.RuneSubString(tr.Text(), 20, 24)
		// 再扔进列表里
		newtermList = append(newtermList, newterm)
		// 计数器会刷到最终次数
		Tcount = index

	})

	// 设置学期计数器
	var I int
	doc.Find("tr[style]").Each(func(index int, tr *goquery.Selection) {
		tr.Find("td[id]").Prev().Each(func(_ int, td *goquery.Selection) {
			// 设置学期名称
			newtermList[I].Term = td.Text()
			// 设置界限, 不超过学期总数
			if I <= Tcount {
				I = I + 1
			}
		})

		tr.Each(func(index int, trr *goquery.Selection) {
			// 新建个成绩结构
			var newscore Score
			// td里面包含具体值
			td := trr.Find("td[width]")
			// 以下对号入座
			newscore.Course = td.First().Text()
			newscore.Type = td.Eq(1).Text()
			newscore.Count = td.Eq(2).Text()
			newscore.Score = td.Eq(3).Text()
			newscore.Credit = td.Eq(4).Text()
			// 过滤掉取了空td的情况
			if newscore.Credit != "" {
				// 将成绩添加到成绩列表
				newtermList[I-1].ScoreList = append(newtermList[I-1].ScoreList, newscore)
			}
			// 加完就删, 避免重复
			tr.Remove().End()

		})
	})

	fmt.Println(newtermList)

}
