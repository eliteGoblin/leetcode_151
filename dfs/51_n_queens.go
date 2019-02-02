package dfs

func solveNQueens(n int) [][]string {
	if n < 4 && 1 != n {
		return [][]string{}
	}
	occupied := make([]int, n)
	for i := range occupied {
		occupied[i] = -1
	}
	res := make([][]string, 0)
	findSolution(&res, occupied, 0)
	return res
}

func findSolution(res *[][]string, occupied []int, row int) {
	if row == len(occupied) {
		(*res) = append(*res, genChessBoard(occupied))
		return
	}
	for col := 0; col < len(occupied); col++ {
		if isValid(occupied, row, col) {
			occupied[row] = col
			findSolution(res, occupied, row+1)
			occupied[row] = -1
		}
	}
}

func isValid(occupied []int, row, col int) bool {
	for i := 0; i <= row; i++ {
		if col == occupied[i] {
			return false
		}
		if abs(row-i) == abs(col-occupied[i]) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func genChessBoard(occupied []int) []string {
	res := make([]string, len(occupied))
	for i := range res {
		for j := 0; j < len(occupied); j++ {
			if j == occupied[i] {
				res[i] += "Q"
			} else {
				res[i] += "."
			}
		}
	}
	return res
}
