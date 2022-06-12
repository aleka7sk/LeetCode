package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = reverseString(words[i])
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	output := ""
	for i := len(s) - 1; i >= 0; i-- {
		output += string(s[i])
	}
	return output
}
