package leetcode

func longestCommonPrefix(strs []string) string {
	temp := strs[0]
	if len(strs) == 1 || temp == "" {
		return temp
	}
	for _, i := range strs[1:] {
		if i == "" {
			return ""
		}
		flag := false
		for index, ch := range i {
			if index+1 <= len(temp) {
				if byte(ch) != temp[index] && index == 0 {
					temp = ""
					return temp
				}
				if byte(ch) != temp[index] {
					temp = i[:index]
					flag = true
					break
				}

			} else {
				if !flag {
					temp = i[:index]
					break
				}
			}
			if !flag && index == len(i)-1 {
				temp = i[:index+1]
			}
		}
	}
	return temp
}
