package leetcode

func isValid(s string) bool {
	open_brackets := []string{"(", "{", "["}
	close_brackets := []string{")", "}", "]"}
	if len(s)%2 == 1 {
		return false
	} else {
		stack := []string{}
		for _, ch := range s {
			if len(stack) == 0 && (ch == ')' || ch == '}' || ch == ']') {
				return false
			} else if ch == '(' || ch == '{' || ch == '[' {
				stack = append(stack, string(ch))
			} else if ch == ')' || ch == '}' || ch == ']' {
				flag := false
				for index, val := range close_brackets {
					if string(ch) == val {
						if stack[len(stack)-1] == open_brackets[index] {
							stack = stack[:len(stack)-1]
							flag = true
							break
						}
					}

				}
				if !flag {
					return false
				}
			}
		}
		if len(stack) == 0 {
			return true
		} else {
			return false
		}
	}
}
