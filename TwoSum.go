package leetcode

func twoSum(nums []int, target int) []int {
	array := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				array[0] = i
				array[1] = j
				return array
			}
		}
	}
	return array
}
