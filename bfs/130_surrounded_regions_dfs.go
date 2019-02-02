package bfs

func solveDFS(board [][]byte) {
	row := len(board)
	if 0 == row {
		return
	}
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if isEdge(board, i, j) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '$':
				board[i][j] = 'O'
			}
		}
	}
}

func isEdge(board [][]byte, i, j int) bool {
	row := len(board)
	col := len(board[0])
	if 0 == i || row-1 == i {
		return true
	}
	if 0 == j || col-1 == j {
		return true
	}
	return false
}

func dfs(board [][]byte, i, j int) {
	row := len(board)
	col := len(board[0])
	if board[i][j] == 'O' {
		board[i][j] = '$'
		if i > 1 && board[i-1][j] == 'O' {
			dfs(board, i-1, j)
		}
		if i < row-2 && board[i+1][j] == 'O' {
			dfs(board, i+1, j)
		}
		if j > 1 && board[i][j-1] == 'O' {
			dfs(board, i, j-1)
		}
		if j < col-2 && board[i][j+1] == 'O' {
			dfs(board, i, j+1)
		}
	}
}
