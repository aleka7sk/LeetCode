package leetcode

func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		if num%2 == 1 {
			sum++
		}
		num = num / 2
	}
	return sum
}
