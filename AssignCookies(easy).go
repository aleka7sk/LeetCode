package leetcode

func findContentChildren(g []int, s []int) int {
	count := 0
	s = SortArray(s)
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if g[i] <= s[j] {
				count++
				s = Pop(s, j)
				break
			}
		}
	}
	return count
}

func SortArray(arr []int) []int {
	for true {
		flag := false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

func Pop(arr []int, position int) []int {
	result := []int{}
	for index, _ := range arr {
		if index == len(arr)-1 {
			return arr[:len(arr)-1]
		} else if index == position {
			result = append(result, arr[:index]...)
			result = append(result, arr[index+1:]...)
			return result
		}
	}
	return arr
}
