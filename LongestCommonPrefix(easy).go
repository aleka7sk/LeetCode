package leetcode

// func longestCommonPrefix(strs []string) string {
// 	temp := strs[0]
// 	if len(strs) == 1 || temp == "" {
// 		return temp
// 	}
// 	for _, i := range strs[1:] {
// 		if i == "" {
// 			return ""
// 		}
// 		flag := false
// 		for index, ch := range i {
// 			if index+1 <= len(temp) {
// 				if byte(ch) != temp[index] && index == 0 {
// 					temp = ""
// 					return temp
// 				}
// 				if byte(ch) != temp[index] {
// 					temp = i[:index]
// 					flag = true
// 					break
// 				}

// 			} else {
// 				if !flag {
// 					temp = i[:index]
// 					break
// 				}
// 			}
// 			if !flag && index == len(i)-1 {
// 				temp = i[:index+1]
// 			}
// 		}
// 	}
// 	return temp
// }

func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}
	var result string
	for i := 0; i < len(strs[0]); i++ {
		flag := true
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i {
				return result
			}
			if strs[0][i] == strs[j][i] {
				continue
			} else {
				flag = false
			}
		}
		if flag {
			result += string(strs[0][i])
		} else {
			return result
		}
	}
	return result
}
