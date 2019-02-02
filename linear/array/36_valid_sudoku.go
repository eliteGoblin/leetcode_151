package array

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if ok := isValidRow(board, i); !ok {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if ok := isValidCol(board, j); !ok {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if ok := isValidSubSudoku(board, i, j); !ok {
				return false
			}
		}
	}
	return true
}

func isValidRow(board [][]byte, row int) bool {
	return isSetValid(board[row])
}

func isValidCol(board [][]byte, col int) bool {
	set := make([]byte, 0)
	for i := 0; i < 9; i++ {
		set = append(set, board[i][col])
	}
	return isSetValid(set)
}

func isValidSubSudoku(board [][]byte, startRow, startCol int) bool {
	set := make([]byte, 0)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			set = append(set, board[i][j])
		}
	}
	return isSetValid(set)
}

func isSetValid(set []byte) bool {
	mp := make(map[byte]bool)
	for i := range set {
		if set[i] == '.' {
			continue
		}
		if set[i] < '1' || set[i] > '9' {
			return false
		}
		if _, ok := mp[set[i]]; ok {
			return false
		}
		mp[set[i]] = true
	}
	return true
}
