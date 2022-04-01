package leetcode

func findSubstring(s string, words []string) []int {
	length := len(words[0])
	result := sortArray(words)
	arr := []int{}
	size := len(words) * length
	for i := 0; i <= len(s)-length*len(words); i++ {
		for j := 0; j < len(words); j++ {
			if s[i:i+length] == words[j] {
				if sortArray(splitString(string(s[i:i+length*len(words)]), length, size)) == result {
					arr = append(arr, i)
					break
				}
			}
		}
	}
	return arr
}

func splitString(str string, length, size int) []string {
	resultarray := []string{}
	for i := 0; i <= size-length; i += length {
		resultarray = append(resultarray, str[i:i+length])
	}
	return resultarray
}

func sortArray(array []string) string {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	res := ""
	for _, i := range array {
		res += i
	}
	return res
}
