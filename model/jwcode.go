package model

type JWCode struct {
	Code  int64
	Value interface{}
}

type TremCodeList struct {
	TremCodeList []JWCode
}
