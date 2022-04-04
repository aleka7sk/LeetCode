package leetcode

func singleNumber(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				flag = true
				nums[j] = 0
				break
			}
		}
		if !flag && nums[i] != 0 {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
