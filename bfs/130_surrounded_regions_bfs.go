package bfs

type Pos struct {
	row int
	col int
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if !isEdge(board, i, j) || board[i][j] != 'O' {
				continue
			}
			q := make(chan Pos, 2*row*col)
			q <- Pos{
				row: i,
				col: j,
			}
			board[i][j] = '$'
			for len(q) > 0 {
				pos := <-q
				if pos.row > 1 && board[pos.row-1][pos.col] == 'O' {
					q <- Pos{
						row: pos.row - 1,
						col: pos.col,
					}
					board[pos.row-1][pos.col] = '$'
				}
				if pos.row < row-2 && board[pos.row+1][pos.col] == 'O' {
					q <- Pos{
						row: pos.row + 1,
						col: pos.col,
					}
					board[pos.row+1][pos.col] = '$'
				}
				if pos.col > 1 && board[pos.row][pos.col-1] == 'O' {
					q <- Pos{
						row: pos.row,
						col: pos.col - 1,
					}
					board[pos.row][pos.col-1] = '$'
				}
				if pos.col < col-2 && board[pos.row][pos.col+1] == 'O' {
					q <- Pos{
						row: pos.row,
						col: pos.col + 1,
					}
					board[pos.row][pos.col+1] = '$'
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch board[i][j] {
			case '$':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}
