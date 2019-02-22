package dp

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	f := make([][]int, len(s))
	for i := range f {
		f[i] = make([]int, len(s))
	}
	maxLen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || f[i+1][j-1] == 1 {
					f[i][j] = 1
					if j-i+1 > maxLen {
						res = s[i : j+1]
					}
				}
			}
		}
	}
	return res
}
