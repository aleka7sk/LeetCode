package leetcode

func letterCombinations(digits string) []string {
	array := make(map[int][]string)
	array[2] = []string{"a", "b", "c"}
	array[3] = []string{"d", "e", "f"}
	array[4] = []string{"g", "h", "i"}
	array[5] = []string{"j", "k", "l"}
	array[6] = []string{"m", "n", "o"}
	array[7] = []string{"p", "q", "r", "s"}
	array[8] = []string{"t", "u", "v"}
	array[9] = []string{"w", "x", "y", "z"}
	arr := []int{}
	for _, elem := range digits {
		arr = append(arr, int(elem-48))
	}
	result := recursion([]string{}, len(arr), 0, array, arr, "")
	return result
}

func recursion(result []string, length int, iter int, dict map[int][]string, arr []int, elem string) []string {
	if length == 0 {
		return result
	}
	for _, char := range dict[arr[iter]] {
		temp := elem + char
		if length == 1 {
			result = append(result, temp)
		}
		result = recursion(result, length-1, iter+1, dict, arr, temp)
	}
	return result
}
