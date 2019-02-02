package dfs

// 题解从 (m-1, n-1) --> (0,0) 巧妙，不用存储upperbound, 直接与0判断，简化代码
// 0 index, 而非1 index, 简化代码，容易理解
func uniquePaths(m int, n int) int {
	if 0 == m || 0 == n {
		return 0
	}
	cache := make([][]int, m) //
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	return uniquePathRecursive(cache, m-1, n-1)
}

func uniquePathRecursive(cache [][]int, curCol, curRow int) int {
	if curCol < 0 || curRow < 0 {
		return 0
	}
	if curCol == 0 && curRow == 0 {
		return 1
	}
	if cache[curCol][curRow] >= 0 {
		return cache[curCol][curRow]
	}
	cache[curCol][curRow] = uniquePathRecursive(cache, curCol-1, curRow) + uniquePathRecursive(cache, curCol, curRow-1)
	return cache[curCol][curRow]
}
