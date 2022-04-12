package leetcode

func repeatedSubstringPattern(s string) bool {
	i := 1
	for i <= len(s)/2 {
		flag := true
		for j := 0; j <= len(s)-2*i; j += i {
			if s[j:j+i] == s[j+i:j+i+i] {
				continue
			} else {
				flag = false
				break
			}
		}
		if flag && len(s)%i == 0 {
			return true
		}
		i++
	}
	return false
}
