package leetcode

func threeSumClosest(nums []int, target int) int {
	sum := nums[0] + nums[1] + nums[2]
	temp := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			for k := 2; k < len(nums); k++ {
				if i != j && i != k && j != k {
					temp = nums[i] + nums[j] + nums[k]
					if abs(temp-target) < abs(sum-target) {
						sum = temp
					}
				}
			}
		}
	}
	return sum
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	} else {
		return num
	}
}
