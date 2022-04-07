package leetcode

// func findDisappearedNumbers(nums []int) []int {
//     set := make(map[int]int)
//     for i := 0; i < len(nums);i++{
//         set[nums[i]] = 1
//     }

//     arr := []int{}
//     for i := 1; i <= len(nums); i++{
//         if set[i] != 1 {
//             arr = append(arr, i)
//         }
//     }
//     return arr
// }

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 && nums[i] != 0 {
			if nums[i] != nums[nums[i]-1] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			} else {
				nums[nums[i]-1] = 0
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
			i--
		}
	}
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			arr = append(arr, i+1)
		}
	}
	return arr
}
