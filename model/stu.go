package model

// Stu 学生信息结构
type Stu struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Faculty string `json:"faculty"`
	Degree  string `json:"degree"`
	EduSys  string `json:"edusys"`
	Major   string `json:"major"`
	Class   string `json:"class"`
	AdmTime string `json:"admtime"`
	GraTime string `json:"gratime"`
}
