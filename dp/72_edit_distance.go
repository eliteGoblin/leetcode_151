package dp

func minDistance(word1 string, word2 string) int {
	if word1 == "" || word2 == "" {
		return max(len(word1), len(word2))
	}
	f := make([][]int, len(word1)+1)
	for i := range f {
		f[i] = make([]int, len(word2)+1)
	}
	for i := range f {
		f[i][0] = i
	}
	for j := range f[0] {
		f[0][j] = j
	}
	for i := 1; i < len(f); i++ {
		for j := 1; j < len(f[0]); j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j-1], f[i][j-1], f[i-1][j]) + 1
			}
		}
	}
	return f[len(f)-1][len(f[0])-1]
}
