package dfs

func solveSudokuPremature(board [][]byte) {
	rowCache := make([]map[byte]bool, 9)
	for i := range rowCache {
		rowCache[i] = make(map[byte]bool)
	}
	colCache := make([]map[byte]bool, 9)
	for i := range colCache {
		colCache[i] = make(map[byte]bool)
	}
	subGridCache := make([]map[byte]bool, 9)
	for i := range subGridCache {
		subGridCache[i] = make(map[byte]bool)
	}
	putOneToBoard := func(row, col int, v byte) {
		board[row][col] = v
		rowCache[row][v] = true
		colCache[col][v] = true
		grid := getSubGrid(row, col)
		subGridCache[grid][v] = true
	}
	unputOneToBoard := func(row, col int, v byte) {
		board[row][col] = '.'
		delete(rowCache[row], v)
		delete(colCache[col], v)
		grid := getSubGrid(row, col)
		delete(subGridCache[grid], v)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v != '.' {
				putOneToBoard(i, j, v)
			}
		}
	}
	isValid := func(row, col int, v byte) bool {
		_, okRow := rowCache[row][v]
		_, okCol := colCache[col][v]
		grid := getSubGrid(row, col)
		_, okGrid := subGridCache[grid][v]
		return !okRow && !okCol && !okGrid
	}
	var dfsSudoku func(i, j int) bool
	dfsSudoku = func(row, col int) bool {
		if row >= 9 || col >= 9 {
			return true
		}
		var v byte
		for v = '1'; v <= '9'; v++ {
			if !isValid(row, col, v) {
				continue
			}
			putOneToBoard(row, col, v)
			i, j := findNextEmptyPos(board, row, col)
			res := dfsSudoku(i, j)
			if res == true {
				return res
			}
			unputOneToBoard(row, col, v)
		}
		return false
	}
	row, col := findNextEmptyPos(board, 0, 0)
	dfsSudoku(row, col)
}

func getSubGrid(i, j int) int {
	row := i / 3
	col := j / 3
	return row*3 + col
}

func findNextEmptyPos(board [][]byte, row, col int) (nextRow, nextCol int) {
	for j := col; j < 9; j++ {
		if board[row][j] == '.' {
			return row, j
		}
	}
	for i := row + 1; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return 9, 9
}
