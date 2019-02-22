package dp

func minCut(s string) int {
	if len(s) <= 1 || len(s) >= 4096 {
		return 0
	}
	isPalimdrome := make([][]bool, len(s))
	for i := range isPalimdrome {
		isPalimdrome[i] = make([]bool, len(s))
	}
	f := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		f[i] = len(s) - i - 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 || isPalimdrome[i+1][j-1] {
					isPalimdrome[i][j] = true
					f[i] = min(f[i], f[j+1]+1)
				}
			}
		}
	}
	return f[0]
}

func min(a ...int) int {
	if len(a) <= 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
