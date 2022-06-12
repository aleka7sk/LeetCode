package leetcode

import "sort"

func sortedSquares(nums []int) []int {
	array := []int{}
	for _, num := range nums {
		array = append(array, num*num)
	}
	sort.Ints(array)
	return array
}
