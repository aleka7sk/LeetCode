package leetcode

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for j := len(nums) - 1; j > 1; j-- {
		if nums[j] < nums[j-1]+nums[j-2] {
			return nums[j] + nums[j-1] + nums[j-2]
		}
	}
	return 0
}
