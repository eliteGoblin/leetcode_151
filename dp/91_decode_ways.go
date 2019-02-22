package dp

func NumDecodingsMemo(s string) int {
	f := make([]int, len(s))
	for i := range f {
		f[i] = -1
	}
	var findDecodeWays func(i int) int
	findDecodeWays = func(i int) int {
		if i >= len(s) {
			return 1
		}
		if f[i] > -1 {
			return f[i]
		}
		if s[i:i+1] == "0" {
			f[i] = 0
			return f[i]
		}
		f[i] = f[i-1]
		if len(s)-i >= 2 {
			f[i] = findDecodeWays(i + 1)
			if s[i:i+2] < "27" {
				f[i] += findDecodeWays(i + 2)
			}
		} else {
			f[i] = 1
		}

		return f[i]
	}
	findDecodeWays(0)
	return f[0]
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	f := make([]int, len(s)+1)
	f[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1:i] == "0" {
			f[i] = 0
		} else {
			f[i] = f[i-1]
			if i > 1 && s[i-2:i] < "27" {
				f[i] += f[i-2]
			}
		}
	}
	return f[len(s)]
}
