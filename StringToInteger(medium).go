package leetcode

func myAtoi(s string) int {
	flag := false
	res := 0
	sign := 1
	for index, char := range s {
		if char == '0' {
			flag = true
		}
		if char == '-' && !flag && res == 0 {
			flag = true
			sign *= -1
			if len(s) >= index+2 {
				if s[index+1] == ' ' {
					return res * sign
				}
			}
			continue
		} else if char == '+' && !flag && res == 0 {
			flag = true
			if len(s) >= index+2 {
				if s[index+1] == ' ' {
					return res * sign
				}
			}
			continue
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-48)
			if len(s) >= index+2 {
				if s[index+1] == ' ' {
					return res * sign
				}
			}
			if res*sign > powNumber(2, 31)-1 {
				return powNumber(2, 31) - 1
			} else if res*sign < powNumber(-2, 31) {
				return powNumber(-2, 31)
			}
		} else if (char < '0' || char > '9') && res != 0 {
			return res * sign
		} else if (char < '0' || char > '9') && res == 0 && char != ' ' {
			return 0
		} else {
			continue
		}
	}
	result := res * sign
	return result

}

func powNumber(num, pow int) int {
	multiply := 1
	for i := 0; i < pow; i++ {
		multiply = multiply * num
	}
	return multiply
}
