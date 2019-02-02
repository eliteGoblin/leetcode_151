package array

func Rotate(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := rows
	// diag
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-i; j++ {
			swap(&matrix[i][j], &matrix[cols-1-j][rows-1-i])
		}
	}
	// up down
	for i := 0; i < rows/2; i++ {
		for j := 0; j < cols; j++ {
			swap(&matrix[i][j], &matrix[rows-1-i][j])
		}
	}
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
