package dp

const (
	INT_MAX = int((^uint(1)) >> 1)
	INT_MIN = -1 * INT_MAX
)

func minPathSumMemo(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var getMinPathSum func(i, j int) int
	getMinPathSum = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return INT_MIN
		}
		if grid[i][j] < 0 {
			return grid[i][j]
		}
		if i == 0 && j == 0 {
			return -1 * grid[0][0]
		}
		grid[i][j] = max(getMinPathSum(i-1, j), getMinPathSum(i, j-1)) - grid[i][j]
		return grid[i][j]
	}
	return getMinPathSum(len(grid)-1, len(grid[0])-1)
}

func minPathSumDP(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
