package model

// Score 成绩结构
type Score struct {
	Course string      `json:"course"`
	Type   interface{} `json:"type"`
	Count  interface{} `json:"count"`
	Score  interface{} `json:"score"`
	Credit interface{} `json:"credit"`
}

// ScoreTerm 学期成绩结构
type ScoreTerm struct {
	Term      string  `json:"term"`
	ScoreList []Score `json:"score_list"`
	AvgScore  string  `json:"avgscore"`
}

// ScoreTermList 学期成绩列表结构
type ScoreTermList []ScoreTerm
