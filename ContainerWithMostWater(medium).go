package leetcode

func maxArea(height []int) int {
	area := 0
	first := 0
	last := len(height) - 1
	for true {
		start := height[first]
		end := height[last]
		var res int
		var minimum int

		if start < end {
			minimum = start
			first++
		} else {
			minimum = end
			last--
		}

		res = minimum * (last - first + 1)

		if area < res {
			area = res
		}

		if first == last {
			break
		}
	}
	return area
}
