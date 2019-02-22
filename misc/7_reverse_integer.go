package misc

func reverse(x int) int {
	if uint(x) == 1<<31 {
		return 0
	}
	neg := false
	var in int64 = int64(x)
	if x < 0 {
		neg = true
		in = -in
	}
	var res int64
	for in > 0 {
		res = res*10 + in%10
		in /= 10
	}
	if res >= 1<<31 {
		return 0
	}
	if neg {
		return -1 * int(res)
	}
	return int(res)
}
