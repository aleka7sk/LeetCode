package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < i+k+1; j++ {
			if j < len(nums) {
				if nums[i] == nums[j] && j-i <= k && j != i {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}
