package leetcode

func countBits(n int) []int {
	array := []int{}
	counter := 0
	for counter <= n {
		i := counter
		count := 0
		for i != 0 {
			if i%2 == 1 {
				count++
			}
			i = i / 2
		}
		array = append(array, count)
		counter++
	}
	return array
}
