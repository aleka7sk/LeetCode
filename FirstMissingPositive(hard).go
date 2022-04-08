package leetcode

// func firstMissingPositive(nums []int) int {
//     set := make(map[int]int)
//     for i := 0; i < len(nums); i++ {
//         set[nums[i]] = 1
//     }

//     for i := 1; i < len(nums) + 1; i++ {
//         _, ok := set[i]
//         if !ok {
//             return i
//         }
//     }
//     return len(nums) + 1
// }

func firstMissingPositive(nums []int) int {
	i := 0
	length := len(nums)
	for i < length {
		if nums[i] > 0 && nums[i] < length {
			j := nums[i] - 1
			if nums[i] != nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				i += 1
			}
		} else {
			i += 1
		}
	}
	for k := 0; k < length; k++ {
		if k+1 != nums[k] {
			return k + 1
		}
	}
	return length + 1
}
