package array

import "math"

func grayCode(n int) []int {
	maxN := int(math.Pow(2, float64(n)))
	res := make([]int, maxN)
	for i := 0; i < maxN; i++ {
		res[i] = i ^ (i / 2)
	}
	return res
}
