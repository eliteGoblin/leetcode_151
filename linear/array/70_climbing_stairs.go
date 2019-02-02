package array

func ClimbStairs(n int) int {
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}
	prev_2 := 1
	prev_1 := 2
	cur := 0
	for i := 3; i <= n; i++ {
		cur = prev_1 + prev_2
		prev_2 = prev_1
		prev_1 = cur
	}
	return cur
}
