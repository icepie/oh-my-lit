package util

import (
	"errors"
	"regexp"
)

// get substring between strings
func GetSubstingBetweenStrings(str string, prefix string, suffix string) (sub string, err error) {

	reg := regexp.MustCompile(prefix + "(.*?)" + suffix)

	result := reg.FindAllStringSubmatch(str, -1)

	if len(result) == 0 {
		return sub, errors.New("no result")
	}

	return result[0][1], nil
}
