package leetcode

func removeElement(nums []int, val int) int {
	count := len(nums)
	for i := 0; i < count; i++ {
		flag := false
		if nums[i] == val {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] != val {
					count = j
					nums[i], nums[j] = nums[j], nums[i]
					flag = true
					break
				}
			}
			if !flag {
				return i
			}
		}
	}
	return count
}
