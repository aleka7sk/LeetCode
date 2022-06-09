package leetcode

func subtractProductAndSum(n int) int {
	sum := 0
	mul := 1
	for n != 0 {
		mod := n % 10
		sum += mod
		mul *= mod
		n = n / 10
	}
	return mul - sum
}
