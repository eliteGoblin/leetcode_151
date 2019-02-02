package brute

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 1)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	res = append(res, []int{nums[0]})
	subSetSizeBeforePrev := 1
	for i := 1; i < len(nums); i++ {
		start := 0
		if nums[i] == nums[i-1] {
			start = subSetSizeBeforePrev
		}
		subSetSizeBeforePrev = len(res)
		for j := range res {
			if j < start {
				continue
			}
			cur := make([]int, len(res[j]))
			copy(cur, res[j])
			cur = append(cur, nums[i])
			res = append(res, cur)
		}
	}
	return res
}
