package leetcode

func maxProfit(prices []int) int {
	var min int
	var max_profit int
	for index := 0; index < len(prices)-1; index++ {
		first := prices[index]
		second := prices[index+1]
		if first < second && index == 0 {
			min = first
		} else if index == 0 {
			min = second
		}
		result := second - first
		if result <= 0 && second < min {
			min = second
		} else if result > 0 {
			if second-min > max_profit {
				max_profit = second - min
			}
		}
	}
	return max_profit
}
