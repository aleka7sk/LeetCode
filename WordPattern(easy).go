package leetcode

func wordPattern(pattern string, s string) bool {
	arr := split(s)
	if len(arr) != len(pattern) {
		return false
	}
	arr_1 := make([]int, len(pattern))
	arr_2 := make([]int, len(pattern))
	count := 0
	count_two := 0
	for i := 0; i < len(pattern); i++ {
		count++
		count_two++
		for j := 0; j < len(pattern); j++ {
			if pattern[i] == pattern[j] {
				arr_1[i] = count
				arr_1[j] = count
			}
			if arr[i] == arr[j] {
				arr_2[i] = count_two
				arr_2[j] = count_two
			}
		}
	}
	for i := 0; i < len(arr_1); i++ {
		if arr_1[i] != arr_2[i] {
			return false
		}
	}
	return true
}

func split(str string) []string {
	arr := []string{}
	temp := ""
	for _, char := range str {
		if char == ' ' {
			arr = append(arr, temp)
			temp = ""
		} else {
			temp += string(char)
		}
	}
	if temp != "" {
		arr = append(arr, temp)
	}
	return arr
}
