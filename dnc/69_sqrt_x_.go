package dnc

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	start, end := 1, x
	for start < end {
		mid := start + (end-start)/2
		if x/mid < mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
