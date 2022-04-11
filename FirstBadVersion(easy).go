package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		guess := isBadVersion(mid)
		if guess == true && !isBadVersion(mid-1) {
			return mid
		} else if guess == false && isBadVersion(mid+1) {
			return mid + 1
		} else if guess == false {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}
