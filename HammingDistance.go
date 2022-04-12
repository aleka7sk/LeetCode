package leetcode

func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		mod_x := x % 2
		mod_y := y % 2
		if mod_x != mod_y {
			count++
		}
		x = x / 2
		y = y / 2
	}
	return count
}
