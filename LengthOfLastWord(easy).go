package leetcode

func lengthOfLastWord(s string) int {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res += string(s[i])
		} else if i < len(s)-1 {
			if s[i] == ' ' && s[i+1] != ' ' {
				break
			}
		}
	}
	return len(res)
}
