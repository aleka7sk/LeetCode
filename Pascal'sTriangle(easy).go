package leetcode

func generate(numRows int) [][]int {
    grid := [][]int{[]int{1}}
    Recursion(numRows, 2, []int{1}, &grid)
    return grid
}

func Recursion(size int, col int, arr []int, grid *[][]int){
    if col > size {
        return 
    }
    array := []int{}
    array = append(array, 1)
    for i := 0; i < len(arr) - 1; i++ {
        array = append(array, arr[i] + arr[i+1])
    }
    array = append(array, 1)
    *grid = append(*grid, array)
    Recursion(size, col + 1, array, grid)
}

