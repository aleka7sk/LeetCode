package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, elem := range nums2 {
		nums1 = append(nums1, elem)
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
	if len(nums1)%2 == 0 {
		return (float64(nums1[len(nums1)/2]) + float64(nums1[len(nums1)/2-1])) / 2
	} else {
		return float64(nums1[len(nums1)/2])
	}
}
