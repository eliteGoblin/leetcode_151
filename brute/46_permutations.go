package brute

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	sub := permute(nums[1:])
	res := make([][]int, 0)
	for _, v := range res {
		for i := 0; i <= len(v); i++ {
			addOne := insert(v, i, nums[0])
			new := clone(addOne)
			res = append(res, new)
		}
	}
	return res
}

func clone(src []int) []int {
	res := make([]int, len(src))
	copy(res, src)
	return res
}

func insert(src []int, pos int, value int) []int {
	s := append(src, 0)
	copy(s[pos+1:], s[pos:])
	s[pos] = value
	return s
}
