package dfs

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if 0 == len(obstacleGrid) {
		return 0
	}
	if 0 == len(obstacleGrid[0]) {
		return 0
	}
	cache := make([][]int, len(obstacleGrid))
	for i := range cache {
		cache[i] = make([]int, len(obstacleGrid[0]))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return uniquePathORecursive(obstacleGrid, len(obstacleGrid)-1, len(obstacleGrid[0])-1, cache)
}

func uniquePathORecursive(obstacleGrid [][]int, i, j int, cache [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if obstacleGrid[i][j] == 1 {
		return 0
	}
	if 0 == j && 0 == i {
		return 1
	}
	if cache[i][j] >= 0 {
		return cache[i][j]
	}
	cache[i][j] = uniquePathORecursive(obstacleGrid, i-1, j, cache) +
		uniquePathORecursive(obstacleGrid, i, j-1, cache)
	return cache[i][j]
}
