package dfs

func totalNQueens(n int) int {
	if n < 4 && 1 != n {
		return 0
	}
	occupied := make([]int, n)
	for i := range occupied {
		occupied[i] = -1
	}
	return findNQueens(occupied, 0)
}

func findNQueens(occupied []int, row int) int {
	if row == len(occupied) {
		return 1
	}
	res := 0
	for col := 0; col < len(occupied); col++ {
		if isValid(occupied, row, col) {
			occupied[row] = col
			res += findNQueens(occupied, row+1)
			occupied[row] = -1
		}
	}
	return res
}
