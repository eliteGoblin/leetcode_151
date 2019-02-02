package dfs

func solveSudoku(board [][]byte) {
	solveSudokuBool(board)
}

func solveSudokuBool(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				var v byte
				for v = '1'; v <= '9'; v++ {
					board[i][j] = v
					if isBoardValid(board, i, j) && solveSudokuBool(board) {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
	}
	return true
}

func isBoardValid(board [][]byte, row, col int) bool {
	target := board[row][col]
	for i := 0; i < 9; i++ {
		if board[i][col] == target && i != row {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[row][j] == target && j != col {
			return false
		}
	}
	for i := 3 * (row / 3); i < 3*((row/3)+1); i++ {
		for j := 3 * (col / 3); j < 3*((col/3)+1); j++ {
			if board[i][j] == target && (i != row) && (j != col) {
				return false
			}
		}
	}
	return true
}
