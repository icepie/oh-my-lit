package model

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// ScoreInfo 成绩信息序列化器
type ScoreInfo struct {
	SI Stu           `json:"stuinfo"`
	TL ScoreTermList `json:"term_list"`
}
