package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] >= 9 && i == 0 {
			digits[i] = 0
			digits = append(digits, 1)
			insert(digits)
		} else if digits[i] == 9 && i != 0 {
			digits[i] = 0
			digits[i-1] += 1
			if digits[i-1] != 10 {
				return digits
			}
		} else if digits[i] == 10 {
			digits[i] = 0
			digits[i-1] += 1
			if digits[i-1] != 10 {
				return digits
			}
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}

func insert(arr []int) {
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}

}
