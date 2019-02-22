package dp

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	for level := len(triangle) - 2; level >= 0; level-- {
		for i := 0; i < len(triangle[level]); i++ {
			minValue := triangle[level+1][i]
			if minValue > triangle[level+1][i+1] {
				minValue = triangle[level+1][i+1]
			}
			triangle[level][i] += minValue
		}
	}
	return triangle[0][0]
}
