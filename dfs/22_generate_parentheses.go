package dfs

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfsParenthesis("(", n-1, n, &res)
	return res
}

func dfsParenthesis(prefix string, leftCt, rightCt int, res *[]string) {
	if leftCt > rightCt {
		return
	}
	if leftCt == 0 && rightCt == 0 {
		*res = append(*res, prefix)
		return
	}
	dfsParenthesis(prefix+"(", leftCt-1, rightCt, res)
	dfsParenthesis(prefix+")", leftCt, rightCt-1, res)
}
