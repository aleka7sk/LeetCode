package leetcode

func getRow(rowIndex int) []int {
	array := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		arr := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 {
				arr = append(arr, 1)
			} else if j == i {
				arr = append(arr, 1)
			} else {
				arr = append(arr, array[i-1][j-1]+array[i-1][j])
			}
		}
		array = append(array, arr)
	}
	return array[rowIndex]
}
