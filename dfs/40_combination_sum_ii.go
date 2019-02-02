package dfs

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if 0 == target {
		return res
	}
	sort.Ints(candidates)
	picked := make([]int, 0, 4096)
	findCombine2(candidates, target, picked, &res)
	return res
}

func findCombine2(candidate []int, target int, picked []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(picked))
		copy(tmp, picked)
		*res = append(*res, tmp)
		return
	}
	pre := -1
	for i, v := range candidate {
		if v == pre {
			continue
		}
		findCombine2(candidate[i+1:], target-v, append(picked, v), res)
		pre = candidate[i]
	}
}
