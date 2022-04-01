package leetcode

func mySqrt(x int) int {
	i := 0
	for true {
		if x >= i*i && x < (i+1)*(i+1) {
			break
		}
		i++
	}
	return i
}
