package dfs

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	got := make([]int, 0)
	sort.Ints(candidates)
	findNumber(target, candidates, got, &res)
	return res
}

func findNumber(target int, candidates []int, got []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		if len(got) > 0 {
			tmp := make([]int, len(got))
			copy(tmp, got)
			*res = append(*res, tmp)
		}
		return
	}
	for i, v := range candidates {
		findNumber(target-v, candidates[i:], append(got, v), res)
	}
}
