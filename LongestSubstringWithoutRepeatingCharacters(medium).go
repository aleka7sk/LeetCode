package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	for i := len(s); i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if j+i <= len(s) {
				if repeatLetter(s[j : j+i]) {
					return len(s[j : j+i])
				}
			} else {
				break
			}
		}
	}
	return 0

}

func repeatLetter(str string) bool {
	if len(str) > 95 {
		return false
	}
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
