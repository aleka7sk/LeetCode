package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	for _, num := range nums1 {
		if _, ok := set[num]; ok {
			set[num]++
		} else {
			set[num] = 1
		}
	}
	array := []int{}
	for _, num := range nums2 {
		if value, ok := set[num]; ok {
			if value > 0 {
				set[num]--
				array = append(array, num)
			}
		}
	}
	return array
}
