package leetcode

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	majority := len(nums) / 2
	dict := make(map[int]int)
	for _, num := range nums {
		if dict[num] >= 0 {
			count := dict[num]
			dict[num] = count + 1
			if count+1 > majority {
				return num
			}
		} else {
			dict[num] = 0
		}
	}
	return 0
}
