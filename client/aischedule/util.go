package aischedule

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func BuildSectionsStr(s []Section) string {
	var sections string
	for _, section := range s {
		sections += strconv.Itoa(section.Section) + ","
	}
	return sections[:len(sections)-1]

}

func BuildStyle() string {
	rand.Seed(time.Now().UnixNano())
	style := StyleList[rand.Intn(len(StyleList))]
	return fmt.Sprintf("{\"color\":\"%s\",\"background\":\"%s\"}", style.Color, style.Background)

}

func BuildWeeksStr(w []int) string {
	var weeks string
	for _, week := range w {
		weeks += strconv.Itoa(week) + ","
	}
	return weeks[:len(weeks)-1]
}

// ScheduleConvertAppCourse 解析出的课表格式转换为导入的格式
func ScheduleConvertAppCourse(schedule Schedule) (courses []AppCourse) {
	for _, course := range schedule.CourseInfos {
		courses = append(courses, AppCourse{
			Day:      int64(course.Day),
			Name:     course.Name,
			Position: course.Position,
			Sections: BuildSectionsStr(course.Sections),
			Style:    BuildStyle(),
			Teacher:  course.Teacher,
			Weeks:    BuildWeeksStr(course.Weeks),
		})
	}
	return
}
