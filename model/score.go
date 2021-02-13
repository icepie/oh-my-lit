package model

// Score 成绩结构
type Score struct {
	Course string      `json:"course"`
	Type   interface{} `json:"type"`
	Count  interface{} `json:"count"`
	Score  interface{} `json:"score"`
	Credit interface{} `json:"credit"`
}

// Term 学期成绩结构
type Term struct {
	Term      string  `json:"term"`
	ScoreList []Score `json:"score_list"`
	AvgScore  string  `json:"avgscore"`
}

// TermList 学期成绩列表结构
type TermList []Term
