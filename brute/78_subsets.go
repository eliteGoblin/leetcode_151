package brute

import "sort"

// recursive
func subsetsRecursive(nums []int) [][]int {
	res := make([][]int, 0)
	getSubset(&res, nums, []int{})
}

func getSubset(res *[][]int, prefix []int, left []int) {
	if len(left) == 0 {
		*res = append(*res, prefix)
		return
	}
	getSubset(res, prefix, left[1:])
	getSubset(res, append(prefix, left[0]), left[1:])
}

// iterative

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for _, v := range nums {
		for _, org := range res {
			cur := make([]int, len(org))
			copy(cur, org)
			cur = append(cur, v)
			res = append(res, cur)
		}
	}
	return res
}
