package leetcode

func intToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	var result string
	for num != 0 {
		if num/value[i] != 0 {
			result += symbol[i]
			num = num - value[i]
		} else {
			if i < 12 {
				i++
			}
		}
	}
	return result
}
