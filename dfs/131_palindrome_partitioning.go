package dfs

func Partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := make([][]string, 0)
	partitionDFS(s[0:], []string{}, &res)
	return res
}

func partitionDFS(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			path = append(path, s[:i])
			partitionDFS(s[i:], path, res)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
