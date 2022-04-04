package leetcode

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	for j := len(s); j >= 0; j-- {
		for i := 0; i < len(s); i++ {
			if i+j <= len(s) {
				if len(s[i:i+j]) == j {
					if reverse(s[i : i+j]) {
						return s[i : i+j]
					}
				}
			} else {
				break
			}
		}
	}
	return ""
}

func reverse(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
