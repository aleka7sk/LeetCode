package leetcode

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = getMaxNum(nums[i], sum+nums[i])
		max = getMaxNum(max, sum)
	}

	return max

}

func getMaxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
