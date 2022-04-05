package leetcode

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		return Itoa(tokens[0])
	}
	result := ""
	for i := 0; i < len(tokens); i++ {
		if isOperator(tokens[i]) {
			count := 2
			arr := []string{}
			for j := i - 1; count != 0; j-- {
				if !isOperator(tokens[j]) && tokens[j] != " " {
					arr = append(arr, tokens[j])
					tokens[j] = " "
					count--
				}
			}
			switch tokens[i] {
			case "*":
				tokens[i] = Atoi(Itoa(arr[0]) * Itoa(arr[1]))
			case "/":
				tokens[i] = Atoi(Itoa(arr[1]) / Itoa(arr[0]))
			case "+":
				tokens[i] = Atoi(Itoa(arr[0]) + Itoa(arr[1]))
			case "-":
				tokens[i] = Atoi(Itoa(arr[1]) - Itoa(arr[0]))
			}
			result = tokens[i]
		}
	}
	return Itoa(result)
}

func Atoi(nbr int) string {
	arr := []string{}
	str := ""
	if nbr < 0 {
		nbr *= -1
		str += "-"
	}
	for nbr != 0 {
		arr = append(arr, string(nbr%10+48))
		nbr = nbr / 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		str += arr[i]
	}
	return str
}

func Itoa(str string) int {
	if str == "" {
		return 0
	}
	number := 0
	sign := 1
	if str[0] == '-' {
		sign *= -1
		str = str[1:]
	}
	for i := 0; i < len(str); i++ {
		{
			number = number*10 + int(str[i]-48)
		}
	}
	return number * sign
}

func isOperator(char string) bool {
	if char == "*" || char == "/" || char == "+" || char == "-" {
		return true
	} else {
		return false
	}
}
