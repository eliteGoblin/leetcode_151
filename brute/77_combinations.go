package brute

// dfs
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	dfs(n, k, 1, []int{}, &res)
	return res
}

func dfs(total int, subSetCount int, curPos int, curSet []int, res *[][]int) {
	if len(curSet) == subSetCount {
		tmp := append([]int{}, curSet...)
		*res = append(*res, tmp)
		return
	}
	for i := curPos; i <= total; i++ {
		dfs(total, subSetCount, i+1, append(curSet, i), res)
	}
}

// recursive
func combineRecursive(n int, k int) [][]int {
	if k > n || k < 0 {
		return [][]int{}
	}
	if 0 == k {
		return [][]int{[]int{}}
	}
	res := combine(n-1, k-1)
	for i := range res {
		res[i] = append(res[i], n)
	}
	other := combine(n-1, k)
	return append(res, other...)
}
