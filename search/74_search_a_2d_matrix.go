package search

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if 0 == row {
		return false
	}
	col := len(matrix[0])
	if 0 == col {
		return false
	}
	i, j := 0, col
	for i <= row && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
