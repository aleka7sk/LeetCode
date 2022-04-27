package leetcode

import (
	"fmt"
	"regexp"
)

func IsMatch(s string, p string) bool {
	change := fmt.Sprintf("^%v$", p)
	res, _ := regexp.Compile(change)

	if res.MatchString(s) {
		return true
	} else {
		return false
	}
}
