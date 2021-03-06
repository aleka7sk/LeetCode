package leetcode

func guessNumber(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := (low + high) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
