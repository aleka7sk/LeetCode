package leetcode

func missingNumber(nums []int) int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = 1
	}
	for i := 0; i <= len(nums); i++ {
		if dict[i] == 1 {
			continue
		} else {
			return i
		}
	}
	return 0
}
