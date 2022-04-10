package leetcode

func sortColors(nums []int) {
	i := 0
	indexOne := 0
	indexTwo := len(nums) - 1
	for i < len(nums) {
		if nums[i] == 2 {
			if i <= indexTwo {
				nums[i], nums[indexTwo] = nums[indexTwo], nums[i]
				indexTwo = indexTwo - 1
			} else {
				break
			}
		} else if nums[i] == 0 {
			if i >= indexOne {
				nums[i], nums[indexOne] = nums[indexOne], nums[i]
				indexOne = indexOne + 1
			} else {
				break
			}
			i++
		} else {
			i++
		}
	}
}
