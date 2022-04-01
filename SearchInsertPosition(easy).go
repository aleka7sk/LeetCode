package leetcode

func searchInsert(nums []int, target int) int {
	for index, elem := range nums {
		if elem >= target {
			return index
		}
	}
	return len(nums)
}
