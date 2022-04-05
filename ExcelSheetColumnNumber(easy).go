package leetcode

func titleToNumber(columnTitle string) int {
	nrune := []rune(columnTitle)
	if len(nrune) == 1 {
		return int(nrune[0] - 64)
	}
	sum := 0
	count := 1
	for i := len(nrune) - 2; i >= 0; i-- {
		sum += int((nrune[i] - 64)) * pow(26, count)
		count++
	}
	sum += int((nrune[len(nrune)-1] - 64))
	return sum
}

func pow(nbr, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * nbr
	}
	return res
}
