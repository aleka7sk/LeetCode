package leetcode

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	sum := 0
	count := 0
	for i := 1; i < len(salary)-1; i++ {
		sum += salary[i]
		count++
	}
	return float64(sum) / float64(count)
}
