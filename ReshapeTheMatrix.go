package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c {
        return mat
    }
    grid := [][]int{}
    array := []int{}
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            array = append(array, mat[i][j])
        }
    }
    count := 0
    for i := 0; i < r; i++ {
        row_array := []int{}
        for j := 0; j < c; j++ {
            row_array = append(row_array, array[count])
            count++
        }
        grid = append(grid, row_array)
    }
    return grid
}unc matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c {
        return mat
    }
    grid := [][]int{}
    array := []int{}
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            array = append(array, mat[i][j])
        }
    }
    count := 0
    for i := 0; i < r; i++ {
        row_array := []int{}
        for j := 0; j < c; j++ {
            row_array = append(row_array, array[count])
            count++
        }
        grid = append(grid, row_array)
    }
    return grid
}
