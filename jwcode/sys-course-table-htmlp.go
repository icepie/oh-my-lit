package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// JWTerm 学期
type JWTerm struct {
	Code interface{}
	Term interface{}
}

// JWSchoolYear 学年
type JWSchoolYear struct {
	Code     interface{}
	Year     interface{}
	TermList []JWTerm
}

// JWClass 班级
type JWClass struct {
	Code  interface{}
	Class interface{}
}

// JWMojaor 专业
type JWMojaor struct {
	Code      interface{}
	Term      interface{}
	ClassList []JWClass
}

// JWFaculty 院系
type JWFaculty struct {
	Code       interface{}
	Faculty    interface{}
	MojaorList []JWMojaor
}

// JWCode 教务系统对应码
type JWCode struct {
	SchoolYearList []JWSchoolYear
	FacultyList    []JWFaculty
	UpdateTime     time.Time
}

func main() {
	f, e := os.Open("jw-sys-term-param.html")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	// 新建教务对应码结构
	var jwcode JWCode

	// 学年学期获取
	{
		var jsy JWSchoolYear
		// 查找可用学年
		doc.Find("select[name=Sel_NJ]").Each(func(index int, st *goquery.Selection) {
			st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
				// 新建可用学年结构变量
				code, _ := op.Attr("value")
				// 代码
				jsy.Code = code
				jsy.Year = op.Text()
				jwcode.SchoolYearList = append(jwcode.SchoolYearList, jsy)
			})

		})

		// 查找可用学期
		doc.Find("select[name=Sel_XNXQ]").Each(func(index int, st *goquery.Selection) {
			st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
				code, _ := op.Attr("value")
				var jt JWTerm

				// 基本赋值
				jt.Code = code
				jt.Term = op.Text()

				// 这俩位置太不一样 所以分开遍历
				for index := range jwcode.SchoolYearList { //获取索引
					if jwcode.SchoolYearList[index].Code == code[0:4] {
						jwcode.SchoolYearList[index].TermList = append(jwcode.SchoolYearList[index].TermList, jt)
					}
				}
			})

		})
	}

	// 查找可用院系
	doc.Find("select[name=Sel_YX]").Each(func(index int, st *goquery.Selection) {
		st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
			class, _ := op.Attr("value")
			fmt.Println(class)
			fmt.Println(op.Text())
		})

	})

	jwcode.UpdateTime = time.Now()

	fmt.Println(jwcode)

}
