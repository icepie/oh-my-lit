package jw

import (
	"bytes"
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
)

// GetJwTime 获取当前教务时间
func (u *JwUser) GetJwTime() (jwTime JwTime, err error) {
	body, err := u.GetBannerRpt()
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return
	}

	// 在线人数
	onstrFull := doc.Find("span#onTimeNum").First().Text()

	// 去掉缩进和空格
	onstrq := strings.Replace(strings.Replace(strings.Replace(onstrFull, "\\t", "", -1), "\n", "", -1), "\t", "", -1)
	// 去掉特殊的符号
	onstrq = strings.Replace(onstrq, "\u00a0", "", -1)
	// 截取一下
	onlineNumstr := exutf8.RuneSubString(onstrq, 5, 10)
	// 尝试转为整形
	onlineNum, _ := strconv.Atoi(onlineNumstr)

	jwTime.OnlineNum = uint(onlineNum)

	rawData := doc.Find("span").Eq(1).Text()

	// 对教务时间进行处理
	// [2021年08月06日 星期五 2021-2022学年第一学期 [假期]]
	jwtimeData := strings.Fields(rawData)

	if len(jwtimeData) < 2 {
		err = errors.New("fatl to get jwtime")
		return
	}

	for i, subStr := range jwtimeData {
		if i != len(jwTime.RawData)-1 {
			jwTime.RawData += subStr + " "
		} else {
			jwTime.RawData += subStr
		}
	}

	YearStr := exutf8.RuneSubString(jwtimeData[2], 0, 4)
	YearNum, _ := strconv.Atoi(YearStr)
	jwTime.Year = uint(YearNum)

	if strings.Contains(jwtimeData[2], "一") {
		jwTime.Term = 0
	}

	if strings.Contains(jwtimeData[2], "二") {
		jwTime.Term = 1
	}

	if strings.Contains(jwtimeData[len(jwtimeData)-1], "假期") {
		jwTime.IsVacation = true
	} else {
		// 周数处理
		weekStr := strings.Trim(jwtimeData[4], "周")
		// 尝试转为整形
		weekNum, _ := strconv.Atoi(weekStr)
		jwTime.Week = uint(weekNum)
	}

	return
}

// GetClassID 获取班级ID
func (u *JwUser) GetClassID(name string) (id string, err error) {

	idMap := make(map[string]string)

	body, err := u.GetClassSelPage()
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(body)))
	if err != nil {
		return
	}

	sel := doc.Find("td#theXZBJ>select")

	sel.Find("option").Each(func(i int, option *goquery.Selection) {
		idMap[option.Text()] = option.AttrOr("value", "")
	})

	id = idMap[name]

	if len(id) == 0 {
		err = errors.New("not found class id")
	}

	return
}

// GetClassID 获取学生ID
func (u *JwUser) GetStuID(name string) (id string, err error) {

	body, err := u.GetListXSRpt(name)
	if err != nil {
		return
	}

	if strings.Contains(body, "@") || !strings.Contains(body, "|") {
		err = errors.New("not found student id")
		return
	}

	id = body[0:12]
	if len(id) == 0 {
		err = errors.New("not found student id")
	}

	return
}
