package leetcode

func reverse(x int) int {
	arr := []int{}
	sign := 1
	result := 0
	if x < 0 {
		sign *= -1
		x *= -1
	}
	for x != 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	for i := 0; i < len(arr); i++ {
		result = result*10 + arr[i]
	}

	if result*sign < powNumber(2, 31)-1 && result*sign > powNumber(-2, 31) {
		return result * sign
	} else {
		return 0
	}
}

func powNumber(num, pow int) int {
	multiply := 1
	for i := 0; i < pow; i++ {
		multiply = multiply * num
	}
	return multiply
}
