package array

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	firstRowContainZero := false
	firstColContainZero := false
	for j := 0; j < cols; j++ {
		if 0 == matrix[0][j] {
			firstRowContainZero = true
			break
		}
	}
	for i := 0; i < rows; i++ {
		if 0 == matrix[i][0] {
			firstColContainZero = true
			break
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if 0 == matrix[i][j] {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for j := 1; j < cols; j++ {
		if 0 == matrix[0][j] {
			for i := 1; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		if 0 == matrix[i][0] {
			for j := 1; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowContainZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColContainZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
