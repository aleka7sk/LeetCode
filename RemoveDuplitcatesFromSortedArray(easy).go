package leetcode

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != -101 {
			nums[index] = nums[i]
			index++
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums[j] = -101
			}
		}
	}
	return index
}
