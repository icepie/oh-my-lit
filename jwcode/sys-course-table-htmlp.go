package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/PuerkitoBio/goquery"
)

// JWTerm 学期
type JWTerm struct {
	Code uint64
	Term string
}

// JWSchoolYear 学年
type JWSchoolYear struct {
	Code     uint64
	Year     string
	TermList []JWTerm
}

// JWClass 班级
type JWClass struct {
	Code  uint64
	Class string
}

// JWMojaor 专业
type JWMojaor struct {
	Code      uint64
	Term      string
	ClassList []JWClass
}

// JWFaculty 院系
type JWFaculty struct {
	Code       uint64
	Faculty    string
	MojaorList []JWMojaor
}

// JWCode 教务系统对应码
type JWCode struct {
	SchoolYearList []JWSchoolYear
	FacultyList    []JWFaculty
	UpdateTime     time.Time
}

// 新建教务对应码结构
var jwcode JWCode

// initConfig 初始化配置
func initConfig(cfpath string) error {

	b, err := yaml.Marshal(jwcode)
	if err != nil {
		return err
	}

	f, err := os.Create(cfpath)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(string(b))

	return nil
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

	// 学年学期获取
	{
		var jsy JWSchoolYear
		// 查找可用学年
		doc.Find("select[name=Sel_NJ]").Each(func(index int, st *goquery.Selection) {
			st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
				// 新建可用学年结构变量
				codestr, _ := op.Attr("value")

				code, err := strconv.ParseUint(codestr, 10, 32)
				if err != nil {
					fmt.Println(err)
				}
				// 代码
				jsy.Code = code
				jsy.Year = op.Text()
				jwcode.SchoolYearList = append(jwcode.SchoolYearList, jsy)
			})

		})

		// 查找可用学期
		doc.Find("select[name=Sel_XNXQ]").Each(func(index int, st *goquery.Selection) {
			st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
				codestr, _ := op.Attr("value")

				code, err := strconv.ParseUint(codestr, 10, 32)
				if err != nil {
					fmt.Println(err)
				}
				var jt JWTerm

				// 基本赋值
				jt.Code = code
				jt.Term = op.Text()

				// 这俩位置太不一样 所以分开遍历
				for index := range jwcode.SchoolYearList { //获取索引
					if jwcode.SchoolYearList[index].Code == code/10 {
						jwcode.SchoolYearList[index].TermList = append(jwcode.SchoolYearList[index].TermList, jt)
					}
				}
			})

		})
	}

	// 查找可用院系
	var jf JWFaculty
	doc.Find("select[name=Sel_YX]").Each(func(index int, st *goquery.Selection) {
		st.Find("option[value]").Each(func(index int, op *goquery.Selection) {
			class, _ := op.Attr("value")
			fmt.Println(class)
			fmt.Println(op.Text())

			// 新建可用学年结构变量
			codestr, _ := op.Attr("value")

			code, err := strconv.ParseUint(codestr, 10, 32)
			if err != nil {
				fmt.Println(err)
			}
			// 代码
			jf.Code = code
			jf.Faculty = op.Text()
			jwcode.FacultyList = append(jwcode.FacultyList, jf)
		})

	})

	jwcode.UpdateTime = time.Now()

	fmt.Println(jwcode)

	initConfig("jwcode.yaml")

}
