package leetcode

import "fmt"

func searchRange(nums []int, target int) []int {
	first := 0
	last := len(nums) - 1
	for first <= last {
		mid := (first + last) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			start := mid
			last := mid
			for i := mid; i > 0; i-- {
				if nums[i-1] == target {
					start = start - 1
				} else {
					break
				}
			}
			for j := mid; j < len(nums)-1; j++ {
				if nums[j+1] == target {
					last = last + 1
				} else {
					break
				}
			}
			return []int{start, last}
		} else if nums[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return []int{-1, -1}
}
