package jw

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// XKJGRptToSchedule 正选结果解析为课程表结构
func XKJGRptToSchedule(body string) (courses []CourseInfo, err error) {

	if strings.Contains(body, "禁止选课！") {
		err = errors.New("Forbidden")
		return
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))

	// 找到那张表
	tbody := doc.Find("table#pageRpt").Find("table").Eq(1).Find("tbody")

	// 开始遍历
	tbody.Find("tr").Each(func(i int, tr *goquery.Selection) {
		_, isExistsClass := tr.Attr("class")
		if isExistsClass {
			// 排除表头
			return
		}

		// 课程名称
		classNameFull := strings.Split(tr.Find("td").Eq(1).Text(), "]")

		// log.Println(classNameFull)

		if len(classNameFull) < 2 {
			return
		}

		classCode := classNameFull[0][1:]
		className := classNameFull[1]

		// 学分
		credit := tr.Find("td").Eq(2).Text()

		// 课程教师
		teacher := strings.Split(tr.Find("td").Eq(4).Text(), "[")[0]

		// 课程地点
		timePlace, _ := tr.Find("td").Eq(10).Html()

		timePlaceList := strings.Split(strings.TrimSuffix(timePlace, "<br/>"), "<br/>")

		// log.Println(timePlaceList)

		for _, tp := range timePlaceList {

			tmp := strings.Split(tp, "/")

			if len(tmp) == 0 {
				continue
			}

			var place string

			if len(tmp) > 1 {
				place = tmp[1]
			}

			if place == "" {
				if strings.HasPrefix(className, "m") {
					place = "MOOC"
				} else if strings.Contains(className, "体育") {
					place = "GYM"
				} else {
					place = "UNKNOWN"
				}
			}

			// 时间
			time := tmp[0]

			time = strings.TrimRight(strings.TrimPrefix(time, "["), "]")
			time = strings.ReplaceAll(strings.ReplaceAll(time, "[", " "), "]", " ")
			time = strings.ReplaceAll(time, "  ", " ")

			timeList := strings.Split(time, " ")

			if len(timeList) < 3 {
				continue
			}

			// log.Println(timeList)

			// log.Println(len(timeList))

			// 定义课程
			var course CourseInfo

			course.Code = classCode
			course.Titile = className
			course.Credit = credit
			course.Teacher = teacher
			course.Location = place
			course.Time = time
			if len(timeList) == 4 {
				course.Weeks = buildWeeks(timeList[2] + timeList[3])
			} else {
				course.Weeks = buildWeeks(timeList[2])
			}
			course.Day = buildDay(timeList[0])

			course.Sections = buildSections(timeList[1])
			course.Start = course.Sections[0]
			course.Duration = len(course.Sections)

			// log.Println(course)

			courses = append(courses, course)

		}

		// tr.Find("td").Each(func(i int, td *goquery.Selection) {
		// 	switch i {
		// 	case 1:
		// 		// 课程名称
		// 		log.Println(td.Text())
		// 	case 2:
		// 		// 学分
		// 		log.Println(td.Text())
		// 	case 4:
		// 		// 教师
		// 		log.Println(td.Text())
		// 	case 10:
		// 		// 时间/地点
		// 		log.Println(td.Text())
		// 	}

		// 	// log.Println(td.Text())
		// })
	})

	return

}

// buildWeeks 构建周数
func buildWeeks(weekStr string) (rte []int) {

	flag := 0

	if strings.Contains(weekStr, "单") {
		flag = 1
	} else if strings.Contains(weekStr, "双") {
		flag = 2
	}

	weekStr = strings.ReplaceAll(strings.ReplaceAll(weekStr, "双", ""), "单", "")
	weekStr = strings.ReplaceAll(weekStr, "周", "")

	weekList := strings.Split(weekStr, ",")
	for _, v := range weekList {
		se := strings.Split(v, "-")
		if len(se) == 2 {
			s, _ := strconv.Atoi(se[0])
			e, _ := strconv.Atoi(se[1])
			for i := s; i <= e; i++ {
				if flag == 1 {
					if i%2 != 0 {
						rte = append(rte, i)
					}
				} else if flag == 2 {
					if i%2 == 0 {
						rte = append(rte, i)
					}
				} else {
					rte = append(rte, i)
				}
			}

		} else {
			i, err := strconv.Atoi(se[0])
			if err == nil {
				rte = append(rte, i)
			}
		}

	}

	return
}

// buildSections 构建节数
func buildSections(sectionStr string) (rte []int) {

	se := strings.Split(strings.ReplaceAll(sectionStr, "节", ""), "-")

	if len(se) == 2 {
		s, _ := strconv.Atoi(se[0])
		e, _ := strconv.Atoi(se[1])
		for i := s; i <= e; i++ {
			rte = append(rte, i)
		}

	} else {
		i, err := strconv.Atoi(se[0])
		if err == nil {
			rte = append(rte, i)
		}
	}

	return

}

// buildDay 构建天
func buildDay(dayStr string) (day int) {
	dayStr = strings.ReplaceAll(strings.ReplaceAll(dayStr, "天", "日"), "星期", "")
	return chDaytoNum(dayStr)
}

// chDaytoNum 根据星期转换为数字
func chDaytoNum(day string) int {
	dayMap := map[string]int{"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7}
	return dayMap[day]
}

// 构建云小洛课表结构
func BuildAirSchedule(raw []CourseInfo) (ret [20][7][5][]CourseInfo) {

	for _, c := range raw {
		for _, week := range c.Weeks {
			var s int
			if c.Start == 1 {
				s = 0
			} else {
				s = (c.Start - 1) / 2
			}

			if c.Day == 0 {
				ret[week-1][6][s] = append(ret[week-1][6][s], c)
			} else {
				ret[week-1][c.Day-1][s] = append(ret[week-1][c.Day-1][s], c)
			}

			if c.Duration == 4 {
				if c.Day == 0 {
					ret[week-1][6][s+2] = append(ret[week-1][6][s+2], c)
				} else {
					ret[week-1][c.Day-1][s] = append(ret[week-1][c.Day-1][s+2], c)
				}
			}

		}
	}
	return ret
}
