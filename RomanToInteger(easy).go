package leetcode

func romanToInt(s string) int {
	sum := 0
	for index := 0; index < len(s); index++ {
		if index < len(s)-1 && getValue(s[index]) < getValue(s[index+1]) {
			sum += getValue(s[index+1]) - getValue(s[index])
			index++
		} else {
			sum += getValue(s[index])
		}
	}
	return sum
}

func getValue(char byte) int {
	switch char {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
