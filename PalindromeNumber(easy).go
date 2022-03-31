package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := []int{}
	for x != 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] == arr[j] {
			continue
		} else {
			return false
		}
	}
	return true
}
