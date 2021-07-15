package util

import (
	"errors"
	"regexp"

	"github.com/axgle/mahonia"
)

// get substring between strings by regexp
func GetSubstringBetweenStringsByRE(str string, prefix string, suffix string) (sub string, err error) {

	reg := regexp.MustCompile(prefix + "(.*?)" + suffix)

	result := reg.FindAllStringSubmatch(str, -1)

	if len(result) == 0 {
		return sub, errors.New("no result")
	}

	return result[0][1], nil
}

// 转回UTF8
func GB18030ToUTF8(s string) string {
	src := mahonia.NewDecoder("gb18030")
	res := src.ConvertString(s)
	tag := mahonia.NewDecoder("utf-8")

	_, cdata, err := tag.Translate([]byte(res), true)
	if err != nil {
		return ""
	}

	result := string(cdata)

	return result
}
