package leetcode

func trap(height []int) int {
	var start int
	var end int
	for i := 0; i < len(height); i++ {
		if height[i] > 0 {
			start = i
			break
		}
	}
	for j := len(height) - 1; j >= 0; j-- {
		if height[j] > 0 {
			end = j
			break
		}
	}
	result := 0
	for start != end {
		max := height[start]
		count := 0
		for i := start + 1; i <= end; i++ {
			if height[i] < max {
				count += max - height[i]
			} else {
				result += count
				max = height[i]
				start = i
				count = 0
			}
		}
		if max == 0 {
			return result
		}
		height[start]--
	}
	return result
}
