package dp

// Recursive也刚能AC
func isInterleaveRecursive(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	var res bool
	if s1[0] == s3[0] {
		if res = isInterleaveRecursive(s1[1:], s2, s3[1:]); res {
			return true
		}
	}
	if s2[0] == s3[0] {
		if res = isInterleaveRecursive(s1, s2[1:], s3[1:]); res {
			return true
		}
	}
	return false
}

// dp
func isInterleaveDP(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	rows := len(s1)
	cols := len(s2)
	// rows * cols 2D array
	f := make([][]bool, rows+1)
	for i := range f {
		f[i] = make([]bool, cols+1)
	}
	f[0][0] = true
	for j := 1; j <= cols; j++ {
		f[0][j] = f[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= rows; i++ {
		f[i][0] = f[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if f[i-1][j] && s1[i-1] == s3[i+j-1] {
				f[i][j] = true
			} else if f[i][j-1] && s2[i-1] == s3[i+j-1] {
				f[i][j] = true
			}
		}
	}
	return f[rows][cols]
}
