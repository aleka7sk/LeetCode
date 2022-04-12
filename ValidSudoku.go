package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			if board[i][j] != '.' {
				if !check(&board, i, j, board[i][j]) {
					return false
				}
			}
		}
	}
	return true
}

func check(board *[][]byte, row, col int, num byte) bool {
	for i := 0; i <= 8; i++ {
		if num == (*board)[row][i] && i != col {
			return false
		}
	}

	for j := 0; j <= 8; j++ {
		if num == (*board)[j][col] && j != row {
			return false
		}
	}

	firstCol := col - col%3
	firstRow := row - row%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[i+firstRow][j+firstCol] == num && (i+firstRow != row || j+firstCol != col) {
				return false
			}
		}
	}
	return true
}
