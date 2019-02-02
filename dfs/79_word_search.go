package dfs

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	flag := initFlagArr(len(board), len(board[0]))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// clearFlagArr(flag) 回溯过程已经revert了
			if ok := existDfs(flag, board, word, i, j); ok {
				return true
			}
		}
	}
	return false
}

func initFlagArr(row, col int) [][]bool {
	flag := make([][]bool, 0)
	for i := 0; i < row; i++ {
		row := make([]bool, col)
		flag = append(flag, row)
	}
	return flag
}

func clearFlagArr(flag [][]bool) {
	for i := range flag {
		for j := range flag {
			flag[i][j] = false
		}
	}
}

func existDfs(flag [][]bool, board [][]byte, word string, row, col int) bool {
	if word == "" {
		return true
	}
	if row < 0 || row >= len(board) {
		return false
	}
	if col < 0 || col >= len(board[0]) {
		return false
	}
	if flag[row][col] {
		return false
	}
	if board[row][col] != byte(word[0]) {
		return false
	}
	flag[row][col] = true
	// left
	res := existDfs(flag, board, word[1:], row, col-1) ||
		existDfs(flag, board, word[1:], row, col+1) ||
		existDfs(flag, board, word[1:], row-1, col) ||
		existDfs(flag, board, word[1:], row+1, col)
	flag[row][col] = false
	return res
}
