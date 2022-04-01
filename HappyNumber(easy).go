package leetcode

func isHappy(n int) bool {
	i := 0
	for true {
		arr := []int{}
		for n != 0 {
			arr = append(arr, n%10)
			n = n / 10
		}
		n = sum(arr)
		if n == 1 {
			return true
		}
		i++
		if i > 10 {
			return false
		}
	}
	return false
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num * num
	}
	return sum
}
