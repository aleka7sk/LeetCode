package leetcode

func twoSum(numbers []int, target int) []int {
	set := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		num := target - numbers[i]
		if value, ok := set[num]; ok {
			return []int{value + 1, i + 1}
		}
		set[numbers[i]] = i
	}
	return []int{}
}
