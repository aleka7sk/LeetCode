package leetcode

func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			flag := false
			if s[i] == s[j] {
				flag = true
			}
			if t[i] == t[j] && !flag {
				return false
			} else if t[i] == t[j] {
				continue
			} else if flag {
				return false
			}
		}
	}
	return true
}
