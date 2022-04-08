package leetcode

func createTargetArray(nums []int, index []int) []int {
	array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var prev = array[index[i]]
		for j := index[i]; j < len(nums)-1; j++ {
			array[j+1], prev = prev, array[j+1]
		}
		array[index[i]] = nums[i]
	}
	return array
}
