package leetcode

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	base := []string{}
	arr := []string{}
	result := ""

	if num < 0 {
		num *= -1
		result += "-"
	}

	for i := '0'; i <= '7'; i++ {
		base = append(base, string(i))
	}

	for num != 0 {
		arr = append(arr, base[num%7])
		num = num / 7
	}

	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i]
	}

	return result
}
