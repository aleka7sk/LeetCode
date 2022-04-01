package leetcode

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	var result string
	for _, ch := range s {
		if ch >= 65 && ch <= 90 {
			result += string(ch + 32)
		} else if ch >= 97 && ch <= 122 || ch >= 48 && ch <= 57 {
			result += string(ch)
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		if result[i] == result[j] {
			continue
		} else {
			return false
		}
	}
	return true
}
