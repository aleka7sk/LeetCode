package leetcode

func convertToTitle(columnNumber int) string {
	arr := []rune{}
	for columnNumber != 0 {
		mod := columnNumber % 26
		if mod != 0 {
			arr = append(arr, rune(mod)+64)
		} else {
			arr = append(arr, rune(26+64))
		}
		if columnNumber != 26 {
			if columnNumber%26 == 0 {
				columnNumber = columnNumber/26 - 1
			} else {
				columnNumber = columnNumber / 26
			}
		} else {
			columnNumber = 0
		}

	}
	var res string
	for i := len(arr) - 1; i >= 0; i-- {
		res += string(arr[i])
	}
	return res
}
