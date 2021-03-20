package model

// StuAccount 学生帐号结构

type StuAccount struct {
	StuID    string `json:"name"`
	PassWord string `json:"password"`
}

// Stu 学生信息结构
type Stu struct {
	Name    string      `json:"name"`
	ID      string      `json:"id"`
	Faculty string      `json:"faculty"`
	Degree  string      `json:"degree"`
	EduSys  interface{} `json:"edusys"`
	Major   string      `json:"major"`
	Class   string      `json:"class"`
	AdmTime string      `json:"admtime"`
	GraTime string      `json:"gratime"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	Token string `json:"token"`
}
