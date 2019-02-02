package stringproblem

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var isPalindrome [1024][1024]bool
	maxLen, left, right := 1, 0, 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if i-j < 2 {
				if s[i] == s[j] {
					isPalindrome[j][i] = true
				}
			} else {
				if s[i] == s[j] && isPalindrome[j+1][i-1] {
					isPalindrome[j][i] = true
				}
			}
			if isPalindrome[j][i] {
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					left = j
					right = i
				}
			}
		}
		isPalindrome[i][i] = true
	}
	return s[left : right+1]
}

func LongestPalindromeMemo(s string) string {
	if "" == s {
		return ""
	}
	var arr [1024][1024]string
	start, end := 0, len(s)-1
	return getLongestPalindromeMemo(&arr, s, start, end)
}

func getLongestPalindromeMemo(arr *[1024][1024]string, s string, start, end int) string {
	if arr[start][end] != "" {
		return arr[start][end]
	}
	if start == end {
		arr[start][start] = s[start : start+1]
	} else if end == start+1 {
		if s[start] == s[end] {
			arr[start][end] = s[start : end+1]
		} else {
			arr[start][end] = s[start:end]
		}
	} else {
		if s[start] == s[end] {
			shortStr := getLongestPalindromeMemo(arr, s, start+1, end-1)
			if len(shortStr) == end-start-1 {
				arr[start][end] = s[start : end+1]
				return arr[start][end]
			}
		}
		str1 := getLongestPalindromeMemo(arr, s, start+1, end)
		str2 := getLongestPalindromeMemo(arr, s, start, end-1)
		if len(str1) > len(str2) {
			arr[start][end] = arr[start+1][end]
		} else {
			arr[start][end] = arr[start][end-1]
		}
	}
	return arr[start][end]
}
