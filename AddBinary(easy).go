package leetcode

func addBinary(a string, b string) string {
	res := ""
	temp := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if temp == 1 && int(a[i]-48) == 0 && int(b[j]-48) == 0 {
			res = "1" + res
			temp = 0
		} else if temp == 1 && int(a[i]-48) == 1 && int(b[j]-48) == 1 {
			res = "1" + res
			temp = 1
		} else if temp == 1 && (int(a[i]-48) == 0 || int(b[j]-48) == 0) || int(a[i]-48) == 1 && int(b[j]-48) == 1 {
			res = "0" + res
			temp = 1
		} else {
			res = string(int(a[i]-48)+int(b[j]-48)+48) + res
		}
	}
	if len(a) > len(b) {
		res, temp = getRemainingPart(len(a)-len(b)-1, temp, a, res)
	} else {
		res, temp = getRemainingPart(len(b)-len(a)-1, temp, b, res)
	}
	if temp == 1 {
		res = "1" + res
	}

	return res
}

func getRemainingPart(length, temp int, str, res string) (string, int) {
	for i := length; i >= 0; i-- {
		if temp == 1 && int(str[i]-48) == 1 {
			res = "0" + res
			temp = 1
		} else if temp == 1 && int(str[i]-48) == 0 {
			res = "1" + res
			temp = 0
		} else {
			res = string(str[i]) + res
		}
	}
	return res, temp
}
