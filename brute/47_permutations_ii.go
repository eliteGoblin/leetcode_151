package brute

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	isEnd := false
	for !isEnd {
		res = append(res, nums)
		nums, isEnd = nextPermutation(nums)
	}
}

func nextPermutation(nums []int) (result []int, isEnd bool) {
	i := len(nums) - 1
	for i > 0 {
		if nums[i-1] < nums[i] {
			break
		}
		i--
	}
	if 0 == i {
		return nums, true
	}
	j := len(nums) - 1
	for j > i {
		if nums[j] > nums[i] {
			break
		}
	}
	res := make([]int, len(nums))
	copy(res, nums)
	res[i], res[j] = res[j], res[i]
	sort.Ints(res[i+1:])
	return res, false
}
