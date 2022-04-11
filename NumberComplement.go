package leetcode

func findComplement(num int) int {
	arr := []int{}
	for num != 0 {
		mod := num % 2
		if mod == 1 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, 1)
		}
		num = num / 2
	}

	number := 0
	for i := 0; i < len(arr); i++ {
		number += arr[i] * Pow(2, i)
	}
	return number
}

func Pow(nbr, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result = result * nbr
	}
	return result
}
