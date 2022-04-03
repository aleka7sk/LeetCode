package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for index, elem := range nums2 {
		nums1[m+index] = elem
	}
	BubbleSort(nums1)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
